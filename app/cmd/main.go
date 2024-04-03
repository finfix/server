package main

import (
	"context"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "server/app/docs"
	"server/app/internal/config"
	accountEndpoint "server/app/internal/services/account/endpoint"
	accountRepository "server/app/internal/services/account/repository"
	accountService "server/app/internal/services/account/service"
	adminEndpoint "server/app/internal/services/admin/endpoint"
	adminRepository "server/app/internal/services/admin/repository"
	adminService "server/app/internal/services/admin/service"
	authEndpoint "server/app/internal/services/auth/endpoint"
	authRepository "server/app/internal/services/auth/repository"
	authService "server/app/internal/services/auth/service"
	"server/app/internal/services/generalRepository"
	"server/app/internal/services/scheduler"
	tgBotService "server/app/internal/services/tgBot/service"
	transactionEndpoint "server/app/internal/services/transaction/endpoint"
	transactionRepository "server/app/internal/services/transaction/repository"
	transactionService "server/app/internal/services/transaction/service"
	userEndpoint "server/app/internal/services/user/endpoint"
	userRepository "server/app/internal/services/user/repository"
	userService "server/app/internal/services/user/service"
	"server/pkg/database"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/middleware"
	"server/pkg/panicRecover"
	"server/pkg/tgBot"
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

	// Передаем в middleware авторизации ключ
	middleware.NewAuthMiddleware(cfg.Token.SigningKey)

	// Подключаемся к базе данных
	logger.Info("Подключаемся к БД")
	db, err := database.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
	defer db.Close()

	// Инициализируем клиента телеграм
	tgBot, tgChat, err := tgBot.Init(cfg.Telegram.Token, cfg.Telegram.ChatID)

	// Регистрируем сервисы
	tgBotService := tgBotService.New(tgBot, tgChat, logger)

	// Регистрируем репозитории
	generalRepository, err := generalRepository.New(db, logger)
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
	accountRepository := accountRepository.New(db, logger)
	transactionRepository := transactionRepository.New(db, logger)
	adminRepository := adminRepository.New(db, logger)
	userRepository := userRepository.New(db, logger)
	authRepository := authRepository.New(db, logger)

	// Регистрируем сервисы
	accountService := accountService.New(
		accountRepository,
		generalRepository,
		transactionRepository,
		userRepository,
		logger,
	)

	transactionService := transactionService.New(
		transactionRepository,
		accountService,
		generalRepository,
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

	scheduler := scheduler.NewScheduler(adminService, logger)
	logger.Info("Start scheduler")
	if err = scheduler.Start(); err != nil {
		logger.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/account", accountEndpoint.NewEndpoint(accountService))
	mux.Handle("/account/", accountEndpoint.NewEndpoint(accountService))
	mux.Handle("/transaction", transactionEndpoint.NewEndpoint(transactionService))
	mux.Handle("/transaction/", transactionEndpoint.NewEndpoint(transactionService))
	mux.Handle("/auth/", authEndpoint.NewEndpoint(authService))
	mux.Handle("/admin/", adminEndpoint.NewEndpoint(adminService))
	mux.Handle("/user/", userEndpoint.NewEndpoint(userService))

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	errs := make(chan error)

	if cfg.HTTP == "" {
		logger.Fatal(errors.InternalServer.New("Переменная окружения LISTEN_HTTP не задана"))
	}
	logger.Info("Server is listening %v", cfg.HTTP)

	go func() {
		server := &http.Server{
			Addr:    cfg.HTTP,
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
