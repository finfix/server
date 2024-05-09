package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
	httpSwagger "github.com/swaggo/http-swagger"

	"server/app/config"
	_ "server/app/docs"
	"server/app/pkg/database"
	"server/app/pkg/errors"
	"server/app/pkg/jwtManager"
	"server/app/pkg/log"
	"server/app/pkg/panicRecover"
	"server/app/pkg/server/middleware"
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
	tagEndpoint "server/app/services/tag/endpoint"
	tagRepository "server/app/services/tag/repository"
	tagService "server/app/services/tag/service"
	transactionEndpoint "server/app/services/transaction/endpoint"
	transactionRepository "server/app/services/transaction/repository"
	transactionService "server/app/services/transaction/service"
	userEndpoint "server/app/services/user/endpoint"
	userRepository "server/app/services/user/repository"
	userService "server/app/services/user/service"
)

// @title COIN Server Documentation
// @version @{version} (build @{build})
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

const version = "@{version}"
const build = "@{build}"

const (
	readHeaderTimeout = 10 * time.Second
)

func main() {
	if err := mainNoExit(); err != nil {
		log.Fatal(context.Background(), err)
	}
}

func mainNoExit() error {

	// Создаем контекст с отменой по вызову функции
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Перехватываем возможную панику
	defer panicRecover.PanicRecover(func(err error) {
		log.Fatal(ctx, err)
	})

	// Парсим флаги
	isSetupTelegram := flag.Bool("telegram", false, "Enabling telegram bot\ntrue:\n\t1. Setup connect\n\t2. Enable sending messages")
	logFormat := flag.String("log-format", string(log.JSONFormat), "text - Human readable string\njson - JSON format")
	envMode := flag.String("env-mode", "local", "Environment mode for log label: test, prod")
	flag.Parse()

	// Инициализируем логгер
	if err := log.Init(
		log.LogFormat(*logFormat),
		map[string]string{
			"env":     *envMode,
			"version": version,
			"build":   build,
		},
	); err != nil {
		return err
	}

	// Получаем конфиг
	log.Info(ctx, "Получаем конфиг")
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	// Инициализируем все сервисы
	log.Info(ctx, "Инициализируем сервисы")
	if err = initServices(cfg); err != nil {
		return err
	}

	// Подключаемся к базе данных
	log.Info(ctx, "Подключаемся к БД")
	db, err := database.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		return err
	}
	defer db.Close()

	// Инициализируем клиента телеграм
	log.Info(ctx, "Инициализируем телеграм клиента")
	tgBot, err := tgBot.NewTgBot(cfg.Telegram.Token, cfg.Telegram.ChatID, *isSetupTelegram)
	if err != nil {
		return err
	}
	defer tgBot.Bot.Close()

	// Регистрируем репозитории
	generalRepository, err := generalRepository.New(db)
	if err != nil {
		return err
	}
	accountRepository := accountRepository.New(db)
	tagRepository := tagRepository.New(db)
	transactionRepository := transactionRepository.New(db)
	settingsRepository := settingsRepository.New(db)
	userRepository := userRepository.New(db)
	authRepository := authRepository.New(db)

	// Регистрируем сервисы
	accountPermisssionsService, err := accountPermisssionsService.New(db)
	if err != nil {
		return err
	}

	settingsService := settingsService.New(
		settingsRepository,
		tgBot,
		settingsService.Version{
			Version: version,
			Build:   build,
		},
		settingsService.Credentials{
			CurrencyProviderAPIKey: cfg.APIKeys.CurrencyProvider,
		},
	)

	accountService := accountService.New(
		accountRepository,
		generalRepository,
		transactionRepository,
		userRepository,
		accountPermisssionsService,
	)

	tagService := tagService.New(
		tagRepository,
		generalRepository,
	)

	transactionService := transactionService.New(
		transactionRepository,
		accountRepository,
		generalRepository,
		accountPermisssionsService,
		tagRepository,
	)

	userService := userService.New(
		userRepository,
		generalRepository,
		accountRepository,

	)

	authService := authService.New(
		authRepository,
		userRepository,
		generalRepository,
		[]byte(cfg.GeneralSalt),

	)

	log.Info(ctx, "Запускаем планировщик")
	if err = scheduler.NewScheduler(settingsService).Start(); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/account", accountEndpoint.NewEndpoint(accountService))
	mux.Handle("/account/", accountEndpoint.NewEndpoint(accountService))
	mux.Handle("/transaction", transactionEndpoint.NewEndpoint(transactionService))
	mux.Handle("/transaction/", transactionEndpoint.NewEndpoint(transactionService))
	mux.Handle("/tag/", tagEndpoint.NewEndpoint(tagService))
	mux.Handle("/tag", tagEndpoint.NewEndpoint(tagService))
	mux.Handle("/auth/", authEndpoint.NewEndpoint(authService))
	mux.Handle("/settings/", settingsEndpoint.NewEndpoint(settingsService, cfg.AdminSecretKey))
	mux.Handle("/user/", userEndpoint.NewEndpoint(userService))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	errs := make(chan error)

	log.Info(ctx, "Запускаем HTTP-сервер")
	if cfg.HTTP == "" {
		return errors.InternalServer.New("Переменная окружения LISTEN_HTTP не задана")
	}
	log.Info(ctx, fmt.Sprintf("Server is listening %v", cfg.HTTP))

	go func() {
		server := &http.Server{
			Addr:                         cfg.HTTP,
			Handler:                      CORS(mux),
			DisableGeneralOptionsHandler: false,
			TLSConfig:                    nil,
			ReadTimeout:                  0,
			ReadHeaderTimeout:            readHeaderTimeout,
			WriteTimeout:                 0,
			IdleTimeout:                  0,
			MaxHeaderBytes:               0,
			TLSNextProto:                 nil,
			ConnState:                    nil,
			ErrorLog:                     nil,
			BaseContext:                  nil,
			ConnContext:                  nil,
		}
		errs <- errors.InternalServer.Wrap(server.ListenAndServe())
	}()

	return <-errs
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
			log.Error(context.Background(), err)
			middleware.DefaultErrorEncoder(context.Background(), w, err)
		})

		handler.ServeHTTP(w, r)
	})
}

func initServices(cfg config.Config) error {

	// Конфигурируем decimal, чтобы в JSON не было кавычек
	decimal.MarshalJSONWithoutQuotes = true

	// Инициализируем JWT-менеджер
	accessTokenTTL, err := time.ParseDuration(cfg.Token.AccessTokenTTL)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	refreshTokenTTL, err := time.ParseDuration(cfg.Token.RefreshTokenTTL)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	jwtManager.Init([]byte(cfg.Token.SigningKey), accessTokenTTL, refreshTokenTTL)

	return nil
}
