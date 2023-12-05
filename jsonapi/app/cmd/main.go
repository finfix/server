package main

import (
	"context"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"logger/app/logging"
	"logger/app/pblogger"
	"pkg/errors"
	"pkg/middleware"
	"pkg/panicRecover"

	"auth/app/proto/pbAuth"
	"core/app/proto/pbAccount"
	"core/app/proto/pbCrontab"
	"core/app/proto/pbTransaction"
	"core/app/proto/pbUser"

	_ "jsonapi/app/docs"
	"jsonapi/app/internal/config"
	"jsonapi/app/internal/services/account"
	"jsonapi/app/internal/services/auth"
	"jsonapi/app/internal/services/crontab"
	"jsonapi/app/internal/services/transaction"
	"jsonapi/app/internal/services/user"

	"google.golang.org/grpc"
)

// @title COIN Server Documentation
// @version 1.0
// @description API Documentation for Coin
// @contact.name Ilia Ivanov
// @contact.email bonavii@icloud.com
// @contact.url

// @securityDefinitions.apikey AuthJWT
// @in header
// @name Authorization
// @description JWT-токен авторизации

// @securityDefinitions.apikey SecretKey
// @in header
// @name MySecretKey
// @description Ключ для доступа к админ-методам

//go:generate go install github.com/swaggo/swag/cmd/swag@v1.8.2
//go:generate go mod download
//go:generate swag init -o ../docs --parseDependency --parseInternal

func main() {

	// Перехватываем панику
	defer panicRecover.PanicRecover(func(err error) {
		logging.GetLogger().Panic(err)
	})

	// Получаем логгер
	logger := logging.GetLogger()

	// Получаем конфиг
	cfg := config.GetConfig()

	// TODO: Уложить все в пару строк
	// Подключаемся к gRPC серверу логгера
	logger.Info("Подключаемся к gRPC-серверу логирования. Address: %v", cfg.Services.Logger.GRPC)
	loggerConn, err := grpc.Dial(cfg.Services.Logger.GRPC, grpc.WithInsecure())
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
	defer loggerConn.Close()
	loggerClient := pblogger.NewLoggerClient(loggerConn)

	// Конфигурируем логгер
	logging.Init(loggerClient, "jsonapi")

	// Передаем в middleware авторизации ключ
	middleware.NewAuthMiddleware(cfg.Token.SigningKey)

	// TODO: Уложить все в пару строк
	// Подключаемся к основному gRPC серверу
	logger.Info("Подключаемся к основному gRPC-серверу. Address: %v", cfg.Services.Core.GRPC)
	coreConn, err := grpc.Dial(cfg.Services.Core.GRPC, grpc.WithInsecure())
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
	defer coreConn.Close()

	// TODO: Уложить все в пару строк
	// Подключаемся к gRPC серверу авторизации
	logger.Info("Подключаемся к gRPC-серверу авторизации. Address: %v", cfg.Services.Auth.GRPC)
	authConn, err := grpc.Dial(cfg.Services.Auth.GRPC, grpc.WithInsecure())
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
	defer authConn.Close()

	mux := http.NewServeMux()
	mux.Handle("/account", account.NewService(pbAccount.NewAccountClient(coreConn)))
	mux.Handle("/account/", account.NewService(pbAccount.NewAccountClient(coreConn)))
	mux.Handle("/transaction", transaction.NewService(pbTransaction.NewTransactionClient(coreConn)))
	mux.Handle("/transaction/", transaction.NewService(pbTransaction.NewTransactionClient(coreConn)))
	mux.Handle("/auth/", auth.NewService(pbAuth.NewAuthClient(authConn)))
	mux.Handle("/crontab/", crontab.NewService(pbCrontab.NewCrontabClient(coreConn)))
	mux.Handle("/user/", user.NewService(pbUser.NewUserClient(coreConn)))

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	errs := make(chan error)

	if cfg.Services.JSON.HTTP == "" {
		logger.Fatal(errors.InternalServer.New("Переменная окружения JSON_LISTEN_HTTP не задана"))
	}
	logger.Info("Server is listening %v", cfg.Services.JSON.HTTP)

	go func() {
		server := &http.Server{
			Addr:    cfg.Services.JSON.HTTP,
			Handler: CORS(mux),
		}
		errs <- errors.InternalServer.Wrap(server.ListenAndServe())
	}()

	logger.Fatal(<-errs)
}

func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			return
		} else {

			// Обрабатываем панику, если она случилась
			defer panicRecover.PanicRecover(func(err error) {
				logging.GetLogger().Panic(err)
				middleware.DefaultErrorEncoder(context.Background(), w, err, func(err error) {})
			})

			h.ServeHTTP(w, r)
		}
	})
}
