package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shopspring/decimal"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/sync/errgroup"

	"server/app/pkg/http/middleware"
	"server/app/pkg/http/router"
	"server/app/pkg/http/server"

	"server/app/config"
	_ "server/app/docs"
	"server/app/pkg/database/postgresql"
	"server/app/pkg/errors"
	"server/app/pkg/jwtManager"
	"server/app/pkg/log"
	"server/app/pkg/log/model"
	"server/app/pkg/migrator"
	"server/app/pkg/panicRecover"
	"server/app/pkg/pushNotificator"
	"server/app/pkg/stackTrace"
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
	"server/migrations"
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
	if err := run(); err != nil {
		log.Fatal(context.Background(), err)
	}
}

func run() error {

	// Основной контекст приложения
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Перехватываем возможную панику
	defer panicRecover.PanicRecover(func(err error) {
		log.Fatal(ctx, err)
	})

	// Парсим флаги
	logFormat := flag.String("log-format", string(log.JSONFormat), "text - Human readable string\njson - JSON format")
	envMode := flag.String("env-mode", "local", "Environment mode for log label: test, prod")
	flag.Parse()

	var logHandlers []log.Handler
	switch *logFormat {
	case "text":
		logHandlers = append(logHandlers, log.NewConsoleHandler(os.Stdout, log.LevelDebug))
	case "json":
		logHandlers = append(logHandlers, log.NewJSONHandler(os.Stdout, log.LevelDebug))
	}

	// Получаем имя хоста
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	// Инициализируем логгер
	log.Init(
		model.SystemInfo{
			Hostname: hostname,
			Version:  version,
			Build:    build,
			Env:      *envMode,
		},
		logHandlers...,
	)

	// Получаем конфиг
	log.Info(ctx, "Получаем конфиг")
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	// Инициализируем все синглтоны
	log.Info(ctx, "Инициализируем синглтоны")
	if err = initSingletons(cfg); err != nil {
		return err
	}

	// Подключаемся к базе данных
	log.Info(ctx, "Подключаемся к БД")
	postrgreSQL, err := postgresql.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		return err
	}
	defer postrgreSQL.Close()

	// Запускаем миграции в базе данных
	// TODO: Подумать, как откатывать миграции при ошибках
	log.Info(ctx, "Запускаем миграции")
	postgreSQLMigrator := migrator.NewMigrator(
		postrgreSQL,
		migrator.MigratorConfig{
			EmbedMigrations: migrations.EmbedMigrationsPostgreSQL,
			Dir:             "pgsql",
		},
	)
	if err = postgreSQLMigrator.Up(ctx); err != nil {
		return err
	}

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
	generalRepository, err := generalRepository.New(postrgreSQL)
	if err != nil {
		return err
	}
	accountGroupRepository := accountGroupRepository.New(postrgreSQL)
	accountRepository := accountRepository.New(postrgreSQL)
	tagRepository := tagRepository.New(postrgreSQL)
	transactionRepository := transactionRepository.New(postrgreSQL)
	settingsRepository := settingsRepository.New(postrgreSQL)
	userRepository := userRepository.New(postrgreSQL)

	// Регистрируем сервисы
	accountPermisssionsService, err := accountPermisssionsService.New(postrgreSQL)
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

	r := router.NewRouter()
	r.Mount("/account", accountEndpoint.NewEndpoint(accountService))
	r.Mount("/accountGroup", accountGroupEndpoint.NewEndpoint(accountGroupService))
	r.Mount("/transaction", transactionEndpoint.NewEndpoint(transactionService))
	r.Mount("/tag", tagEndpoint.NewEndpoint(tagService))
	r.Mount("/auth", authEndpoint.NewEndpoint(authService))
	r.Mount("/settings", settingsEndpoint.NewEndpoint(settingsService))
	r.Mount("/user", userEndpoint.NewEndpoint(userService))
	r.Mount("/swagger", httpSwagger.WrapHandler)

	server, err := server.GetDefaultServer(cfg.HTTP, r)
	if err != nil {
		return err
	}

	// Создаем wait группу
	eg, ctx := errgroup.WithContext(ctx)

	// Запускаем HTTP-сервер
	eg.Go(func() error { return server.Serve(ctx) })

	// Запускаем горутину, ожидающую завершение контекста
	eg.Go(func() error {

		// Если контекст завершился, значит процесс убили
		<-ctx.Done()

		// Плавно завершаем работу сервера
		server.Shutdown(ctx)

		return nil
	})

	// Ждем завершения контекста или ошибок в горутинах
	return eg.Wait()
}

func initSingletons(cfg config.Config) error {

	stackTrace.Init(cfg.ServiceName)

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

	if err = middleware.Init(cfg.ServiceName); err != nil {
		return err
	}

	return nil
}
