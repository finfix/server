package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/shopspring/decimal"
	httpSwagger "github.com/swaggo/http-swagger"

	"server/app/config"
	_ "server/app/docs"
	"server/app/pkg/database"
	"server/app/pkg/errors"
	"server/app/pkg/jwtManager"
	"server/app/pkg/log"
	"server/app/pkg/panicRecover"
	"server/app/pkg/pushNotificator"
	"server/app/pkg/tgBot"
	accountEndpoint "server/app/services/account/endpoint"
	accountRepository "server/app/services/account/repository"
	accountService "server/app/services/account/service"
	accountGroupEndpoint "server/app/services/accountGroup/endpoint"
	accountGroupRepository "server/app/services/accountGroup/repository"
	accountGroupService "server/app/services/accountGroup/service"
	accountPermisssionsService "server/app/services/accountPermissions"
	authEndpoint "server/app/services/auth/endpoint"
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
	tgBot, err := tgBot.NewTgBot(cfg.Telegram.Token, cfg.Telegram.ChatID, cfg.Telegram.Enabled)
	if err != nil {
		return err
	}
	if cfg.Telegram.Enabled {
		defer tgBot.Bot.Close()
	}

	log.Info(ctx, "Инициализируем пуши")
	pushNotificator, err := pushNotificator.NewPushNotificator(cfg.Notifications.Enabled, pushNotificator.APNsCredentials{
		TeamID:      cfg.Notifications.APNs.TeamID,
		KeyID:       cfg.Notifications.APNs.KeyID,
		KeyFilePath: cfg.Notifications.APNs.KeyFilePath,
	})
	if err != nil {
		return err
	}

	// Регистрируем репозитории
	generalRepository, err := generalRepository.New(db)
	if err != nil {
		return err
	}
	accountGroupRepository := accountGroupRepository.New(db)
	accountRepository := accountRepository.New(db)
	tagRepository := tagRepository.New(db)
	transactionRepository := transactionRepository.New(db)
	settingsRepository := settingsRepository.New(db)
	userRepository := userRepository.New(db)

	// Регистрируем сервисы
	accountPermisssionsService, err := accountPermisssionsService.New(db)
	if err != nil {
		return err
	}

	accountGroupService := accountGroupService.New(
		accountGroupRepository,
		generalRepository,
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
		pushNotificator,
		[]byte(cfg.GeneralSalt),
	)

	settingsService := settingsService.New(
		settingsRepository,
		userService,
		tgBot,
		settingsService.Version{
			Version: version,
			Build:   build,
		},
		settingsService.Credentials{
			CurrencyProviderAPIKey: cfg.APIKeys.CurrencyProvider,
		},
	)

	authService := authService.New(
		userRepository,
		generalRepository,
		[]byte(cfg.GeneralSalt),
	)

	log.Info(ctx, "Запускаем планировщик")
	if err = scheduler.NewScheduler(settingsService).Start(); err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Mount("/account", accountEndpoint.NewEndpoint(accountService))
	r.Mount("/accountGroup", accountGroupEndpoint.NewEndpoint(accountGroupService))
	r.Mount("/transaction", transactionEndpoint.NewEndpoint(transactionService))
	r.Mount("/tag", tagEndpoint.NewEndpoint(tagService))
	r.Mount("/auth", authEndpoint.NewEndpoint(authService))
	r.Mount("/settings", settingsEndpoint.NewEndpoint(settingsService))
	r.Mount("/user", userEndpoint.NewEndpoint(userService))
	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte("OK")) })
	r.Mount("/swagger", httpSwagger.WrapHandler)

	errs := make(chan error)

	log.Info(ctx, "Запускаем HTTP-сервер")
	if cfg.HTTP == "" {
		return errors.InternalServer.New("Переменная окружения LISTEN_HTTP не задана")
	}
	log.Info(ctx, fmt.Sprintf("Server is listening %v", cfg.HTTP))

	go func() {
		server := &http.Server{
			Addr:                         cfg.HTTP,
			Handler:                      r,
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
