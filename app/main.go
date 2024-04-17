package main

import (
	"context"
	"net/http"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"

	"server/app/config"
	_ "server/app/docs"
	"server/app/pkg/database"
	"server/app/pkg/errors"
	"server/app/pkg/logging"
	"server/app/pkg/middleware"
	"server/app/pkg/panicRecover"
	"server/app/pkg/tgBot"
	accountEndpoint "server/app/services/account/endpoint"
	accountRepository "server/app/services/account/repository"
	accountService "server/app/services/account/service"
	accountPermisssionsService "server/app/services/accountPermissions"
	adminEndpoint "server/app/services/admin/endpoint"
	adminRepository "server/app/services/admin/repository"
	adminService "server/app/services/admin/service"
	authEndpoint "server/app/services/auth/endpoint"
	authRepository "server/app/services/auth/repository"
	authService "server/app/services/auth/service"
	"server/app/services/generalRepository"
	"server/app/services/scheduler"
	tgBotService "server/app/services/tgBot/service"
	transactionEndpoint "server/app/services/transaction/endpoint"
	transactionRepository "server/app/services/transaction/repository"
	transactionService "server/app/services/transaction/service"
	userEndpoint "server/app/services/user/endpoint"
	userRepository "server/app/services/user/repository"
	userService "server/app/services/user/service"
)

// @title COIN Server Documentation
// @version 1.0.3 (build 10)
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
// @name AdminSecretKey
// @description Ключ для доступа к админ-методам

//go:generate go install github.com/swaggo/swag/cmd/swag@v1.8.2
//go:generate go mod download
//go:generate swag init -o docs --parseInternal

const version = "1.0.3"
const build = "10"

const (
	readHeaderTimeout = 10 * time.Second
)

var erasePathOption = errors.Options{ErasePath: true}

func main() {

	// Перехватываем панику
	defer panicRecover.PanicRecover(func(err error) {
		logging.GetLogger().Panic(err)
	})

	// Получаем логгер
	logger := logging.GetLogger()

	// Получаем конфиг
	cfg := config.GetConfig()

	// Передаем в middleware авторизации ключ
	middleware.NewAuthMiddleware(cfg.Token.SigningKey)

	// Подключаемся к базе данных
	logger.Info("Подключаемся к БД")
	db, err := database.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err, erasePathOption))
	}
	defer db.Close()

	// Инициализируем клиента телеграм
	logger.Info("Инициализируем телеграм клиента")
	tgBot, tgChat, err := tgBot.Init(cfg.Telegram.Token, cfg.Telegram.ChatID)
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err, erasePathOption))
	}

	// Регистрируем сервисы
	tgBotService := tgBotService.New(tgBot, tgChat, logger)

	// Регистрируем репозитории
	generalRepository, err := generalRepository.New(db, logger)
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err, erasePathOption))
	}
	accountRepository := accountRepository.New(db, logger)
	transactionRepository := transactionRepository.New(db, logger)
	adminRepository := adminRepository.New(db, logger)
	userRepository := userRepository.New(db, logger)
	authRepository := authRepository.New(db, logger)

	// Регистрируем сервисы
	accountPermisssionsService, err := accountPermisssionsService.New(
		db,
		logger,
	)
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err, erasePathOption))
	}

	accountService := accountService.New(
		accountRepository,
		generalRepository,
		transactionRepository,
		userRepository,
		accountPermisssionsService,
		logger,
	)

	transactionService := transactionService.New(
		transactionRepository,
		accountRepository,
		generalRepository,
		accountPermisssionsService,
		logger,
	)

	adminService := adminService.New(
		adminRepository,
		tgBotService,
		logger,
	)

	userService := userService.New(
		userRepository,
		generalRepository,
		accountRepository,
		logger,
	)

	authService := authService.New(
		authRepository,
		userRepository,
		logger,
	)

	logger.Info("Запускаем планировщик")
	if err = scheduler.NewScheduler(adminService, logger).Start(); err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err, erasePathOption))
	}

	mux := http.NewServeMux()
	mux.Handle("/account", accountEndpoint.NewEndpoint(accountService))
	mux.Handle("/account/", accountEndpoint.NewEndpoint(accountService))
	mux.Handle("/transaction", transactionEndpoint.NewEndpoint(transactionService))
	mux.Handle("/transaction/", transactionEndpoint.NewEndpoint(transactionService))
	mux.Handle("/auth/", authEndpoint.NewEndpoint(authService))
	mux.Handle("/admin/", adminEndpoint.NewEndpoint(adminService))
	mux.Handle("/user/", userEndpoint.NewEndpoint(userService))

	mux.HandleFunc("/version", getVersionHandleFunc(version, build))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	errs := make(chan error)

	logger.Info("Запускаем HTTP-сервер")
	if cfg.HTTP == "" {
		logger.Fatal(errors.InternalServer.New("Переменная окружения LISTEN_HTTP не задана"))
	}
	logger.Info("Server is listening %v", cfg.HTTP)

	go func() {
		server := &http.Server{
			Addr:              cfg.HTTP,
			Handler:           CORS(mux),
			ReadHeaderTimeout: readHeaderTimeout,
		}
		errs <- errors.InternalServer.Wrap(server.ListenAndServe())
	}()

	logger.Fatal(<-errs)
}

func CORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			return
		}

		// Обрабатываем панику, если она случилась
		defer panicRecover.PanicRecover(func(err error) {
			logging.GetLogger().Panic(err)
			middleware.DefaultErrorEncoder(context.Background(), w, err)
		})

		handler.ServeHTTP(w, r)
	})
}

func getVersionHandleFunc(version, build string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		versionResponse := struct {
			Version string `json:"version"`
			Build   string `json:"build"`
		}{
			Version: version,
			Build:   build,
		}
		w.Header().Set("Content-Type", "application/json")
		_ = middleware.DefaultResponseEncoder(context.Background(), w, versionResponse)
	}
}
