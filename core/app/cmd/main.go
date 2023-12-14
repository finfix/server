package main

import (
	accountEndpoint "core/app/internal/services/account/endpoints"
	accountRepository "core/app/internal/services/account/repository"
	accountService "core/app/internal/services/account/service"
	crontabEndpoint "core/app/internal/services/crontab/endpoints"
	crontabRepository "core/app/internal/services/crontab/repository"
	crontabService "core/app/internal/services/crontab/service"
	transactionEndpoint "core/app/internal/services/transaction/endpoints"
	transactionRepository "core/app/internal/services/transaction/repository"
	transactionService "core/app/internal/services/transaction/service"
	userEndpoint "core/app/internal/services/user/endpoints"
	userRepository "core/app/internal/services/user/repository"
	userService "core/app/internal/services/user/service"
	"core/app/proto/pbUser"

	"net"

	"google.golang.org/grpc"

	"logger/app/logging"
	loggingMiddleware "logger/app/logging/middleware"
	"logger/app/pblogger"
	"pkg/database"
	"pkg/errors"
	"pkg/panicRecover"

	"core/app/internal/config"
	"core/app/internal/services/generalRepository"
	"core/app/proto/pbAccount"
	"core/app/proto/pbCrontab"
	"core/app/proto/pbTransaction"
)

// @title COIN Server Documentation
// @version 1.0
// @description API Documentation for Coin
// @contact.name Ilia Ivanov
// @contact.email bonavii@icloud.com
// @contact.url https://gitlab.com/myCoin

func main() {

	// Логируем возможную панику
	defer panicRecover.PanicRecover(func(err error) {
		logging.GetLogger().Panic(err)
	})

	// Получаем логгер
	logger := logging.GetLogger()

	// Получаем конфиг
	cfg := config.GetConfig()

	// Подключаемся по gRPC к сервису логирования
	logger.Info("Подключаемся к gRPC-серверу логирования. Address: %v", cfg.Services.Logger.GRPC)
	conn, err := grpc.Dial(cfg.Services.Logger.GRPC, grpc.WithInsecure())
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
	defer conn.Close()

	// Инициализируем клиента сервиса логирования
	loggerClient := pblogger.NewLoggerClient(conn)

	// Даем логеру доступ к сервису логирования
	logging.Init(loggerClient, "core")

	// Подключаемся к базе данных
	logger.Info("Подключаемся к БД")
	db, err := database.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
	defer db.Close()

	// Проверяем наличие порта для запуска gRPC-сервера
	if cfg.Services.Core.GRPC == "" {
		logger.Fatal(errors.InternalServer.New("Переменная окружения CORE_LISTEN_GRPC не задана"))
	}

	// Начинаем слушать порт
	lis, err := net.Listen("tcp", cfg.Services.Core.GRPC)
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}

	// Регистрируем gRPC-сервер
	s := grpc.NewServer(grpc.UnaryInterceptor(loggingMiddleware.LoggingError))

	// Регистрируем репозитории
	generalRepository, err := generalRepository.New(db, logger)
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
	accountRepository := accountRepository.New(db, logger)
	transactionRepository := transactionRepository.New(db, logger)
	crontabRepository := crontabRepository.New(db, logger)
	userRepository := userRepository.New(db, logger)

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

	crontabService := crontabService.New(
		crontabRepository,
		logger,
	)

	userService := userService.New(
		userRepository,
		generalRepository,
		accountRepository,
		logger,
	)

	// Регистрируем эндпоинты
	pbAccount.RegisterAccountServer(s, accountEndpoint.New(accountService, logger))
	pbTransaction.RegisterTransactionServer(s, transactionEndpoint.New(transactionService, logger))
	pbUser.RegisterUserServer(s, userEndpoint.New(userService, logger))
	pbCrontab.RegisterCrontabServer(s, crontabEndpoint.New(crontabService, logger))

	// Запускаем  gRPC-сервер
	logger.Info("gRPC-server is listening %v", cfg.Services.Core.GRPC)
	if err := s.Serve(lis); err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
}
