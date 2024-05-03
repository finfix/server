package main

import (
	"context"
	"flag"
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
// @version 1.1.0 (build 13)
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

const version = "1.1.0"
const build = "13"

const (
	readHeaderTimeout = 10 * time.Second
)

func main() {

	mainCtx := context.Background()

	// Перехватываем панику
	defer panicRecover.PanicRecover(func(err error) {
		log.Panic(mainCtx, err)
	})

	isSetupTelegram := false
	flag.BoolVar(&isSetupTelegram, "telegram", false, "Включаем ли телеграм бот")
	flag.Parse()

	// Получаем конфиг
	cfg := config.GetConfig()

	// Инициализируем все сервисы
	Init(cfg)

	// Подключаемся к базе данных
	log.Info(mainCtx, "Подключаемся к БД")
	db, err := database.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Инициализируем клиента телеграм
	log.Info(mainCtx, "Инициализируем телеграм клиента")
	tgBot, err := tgBot.NewTgBot(cfg.Telegram.Token, cfg.Telegram.ChatID, isSetupTelegram)
	if err != nil {
		log.Fatal(err)
	}
	defer tgBot.Bot.Close()

	// Регистрируем репозитории
	generalRepository, err := generalRepository.New(db)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	settingsService := settingsService.New(
		settingsRepository,
		tgBot,
		version,
		build,
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

	log.Info(mainCtx, "Запускаем планировщик")
	if err = scheduler.NewScheduler(settingsService).Start(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/account", accountEndpoint.NewEndpoint(accountService))
	mux.Handle("/account/", accountEndpoint.NewEndpoint(accountService))
	mux.Handle("/transaction", transactionEndpoint.NewEndpoint(transactionService))
	mux.Handle("/transaction/", transactionEndpoint.NewEndpoint(transactionService))
	mux.Handle("/tag/", tagEndpoint.NewEndpoint(tagService))
	mux.Handle("/tag", tagEndpoint.NewEndpoint(tagService))
	mux.Handle("/auth/", authEndpoint.NewEndpoint(authService))
	mux.Handle("/settings/", settingsEndpoint.NewEndpoint(settingsService))
	mux.Handle("/user/", userEndpoint.NewEndpoint(userService))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	errs := make(chan error)

	log.Info(mainCtx, "Запускаем HTTP-сервер")
	if cfg.HTTP == "" {
		log.Fatal(errors.InternalServer.New("Переменная окружения LISTEN_HTTP не задана"))
	}
	log.Info(mainCtx, "Server is listening %v", cfg.HTTP)

	go func() {
		server := &http.Server{
			Addr:              cfg.HTTP,
			Handler:           CORS(mux),
			ReadHeaderTimeout: readHeaderTimeout,
		}
		errs <- errors.InternalServer.Wrap(server.ListenAndServe())
	}()

	log.Fatal(<-errs)
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
			log.Panic(context.Background(), err)
			middleware.DefaultErrorEncoder(context.Background(), w, err)
		})

		handler.ServeHTTP(w, r)
	})
}

func Init(cfg *config.Config) error {

	// Конфигурируем decimal, чтобы в JSON не было кавычек
	decimal.MarshalJSONWithoutQuotes = true

	// Инициализируем JWT-менеджер
	accessTokenTTL, err := time.ParseDuration(cfg.Token.AccessTokenTTL)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	jwtManager.Init([]byte(cfg.Token.SigningKey), accessTokenTTL)

	return nil
}
