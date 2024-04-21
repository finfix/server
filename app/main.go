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
	authEndpoint "server/app/services/auth/endpoint"
	authRepository "server/app/services/auth/repository"
	authService "server/app/services/auth/service"
	"server/app/services/generalRepository"
	"server/app/services/scheduler"
	settingsEndpoint "server/app/services/settings/endpoint"
	settingsRepository "server/app/services/settings/repository"
	settingsService "server/app/services/settings/service"
	tgBotService "server/app/services/tgBot/service"
	transactionEndpoint "server/app/services/transaction/endpoint"
	transactionRepository "server/app/services/transaction/repository"
	transactionService "server/app/services/transaction/service"
	userEndpoint "server/app/services/user/endpoint"
	userRepository "server/app/services/user/repository"
	userService "server/app/services/user/service"
)

// @title COIN Server Documentation
// @version 1.0.3 (build 11)
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
const build = "11"

const (
	readHeaderTimeout = 10 * time.Second
)

func main() {

	mainCtx := context.Background()

	// Перехватываем панику
	defer panicRecover.PanicRecover(func(err error) {
		logging.GetLogger().Panic(mainCtx, err)
	})

	// Получаем логгер
	logger := logging.GetLogger()

	// Получаем конфиг
	cfg := config.GetConfig()

	// Передаем в middleware авторизации ключ
	middleware.NewAuthMiddleware(cfg.Token.SigningKey)

	// Подключаемся к базе данных
	logger.Info(mainCtx, "Подключаемся к БД")
	db, err := database.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// Инициализируем клиента телеграм
	logger.Info(mainCtx, "Инициализируем телеграм клиента")
	tgBot, tgChat, err := tgBot.Init(cfg.Telegram.Token, cfg.Telegram.ChatID)
	if err != nil {
		logger.Fatal(err)
	}

	// Регистрируем сервисы
	tgBotService := tgBotService.New(tgBot, tgChat, logger)

	// Регистрируем репозитории
	generalRepository, err := generalRepository.New(db, logger)
	if err != nil {
		logger.Fatal(err)
	}
	accountRepository := accountRepository.New(db, logger)
	transactionRepository := transactionRepository.New(db, logger)
	settingsRepository := settingsRepository.New(db, logger)
	userRepository := userRepository.New(db, logger)
	authRepository := authRepository.New(db, logger)

	// Регистрируем сервисы
	accountPermisssionsService, err := accountPermisssionsService.New(
		db,
		logger,
	)
	if err != nil {
		logger.Fatal(err)
	}

	settingsService := settingsService.New(
		settingsRepository,
		tgBotService,
		logger,
		version,
		build,
	)

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

	logger.Info(mainCtx, "Запускаем планировщик")
	if err = scheduler.NewScheduler(settingsService, logger).Start(); err != nil {
		logger.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/account", accountEndpoint.NewEndpoint(logger, accountService))
	mux.Handle("/account/", accountEndpoint.NewEndpoint(logger, accountService))
	mux.Handle("/transaction", transactionEndpoint.NewEndpoint(logger, transactionService))
	mux.Handle("/transaction/", transactionEndpoint.NewEndpoint(logger, transactionService))
	mux.Handle("/tag/", tagEndpoint.NewEndpoint(logger, tagService))
	mux.Handle("/tag", tagEndpoint.NewEndpoint(logger, tagService))
	mux.Handle("/auth/", authEndpoint.NewEndpoint(logger, authService))
	mux.Handle("/settings/", settingsEndpoint.NewEndpoint(logger, settingsService))
	mux.Handle("/user/", userEndpoint.NewEndpoint(logger, userService))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	errs := make(chan error)

	logger.Info(mainCtx, "Запускаем HTTP-сервер")
	if cfg.HTTP == "" {
		logger.Fatal(errors.InternalServer.New("Переменная окружения LISTEN_HTTP не задана"))
	}
	logger.Info(mainCtx, "Server is listening %v", cfg.HTTP)

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
			logging.GetLogger().Panic(context.Background(), err)
			middleware.DefaultErrorEncoder(context.Background(), w, err)
		})

		handler.ServeHTTP(w, r)
	})
}
