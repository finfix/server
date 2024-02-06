package main

import (
	authEndpoint "auth/app/internal/services/auth/endpoints"
	authRepo "auth/app/internal/services/auth/repository"
	authService "auth/app/internal/services/auth/service"
	userService "auth/app/internal/services/user"
	pbUser "core/app/proto/pbUser"
	grpcPkg "pkg/grpc"

	"google.golang.org/grpc"

	"logger/app/logging"
	loggingMiddleware "logger/app/logging/middleware"
	"logger/app/pblogger"
	"pkg/database"
	"pkg/errors"
	"pkg/panicRecover"

	"auth/app/internal/config"
	"auth/app/proto/pbAuth"
)

func main() {

	// Логируем возможную панику
	defer panicRecover.PanicRecover(func(err error) {
		logging.GetLogger().Panic(err)
	})

	// Получаем логгер
	logger := logging.GetLogger()

	// Получаем конфиг
	cfg := config.GetConfig()

	// TODO: Все это дело как-то в пару строк уместить
	// Подключаемся по gRPC к сервису логирования
	logger.Info("Подключаемся к gRPC-серверу логирования. Address: %v", cfg.Services.Logger.GRPC)
	_loggerConn, err := grpc.Dial(cfg.Services.Logger.GRPC, grpc.WithInsecure())
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}

	defer _loggerConn.Close()

	// Инициализируем клиента сервиса логирования
	loggerClient := pblogger.NewLoggerClient(_loggerConn)

	// Даем логеру доступ к сервису логирования
	logging.Init(loggerClient, "auth")

	// TODO: Все это дело как-то в пару строк уместить
	// Подключаемся по gRPC к сервису core
	logger.Info("Подключаемся к gRPC-серверу core. Address: %v", cfg.Services.Core.GRPC)
	_userConn, err := grpc.Dial(cfg.Services.Core.GRPC, grpc.WithInsecure())
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
	defer _userConn.Close()

	// Инициализируем клиента сервиса логирования
	userConn := pbUser.NewUserClient(_userConn)

	// Подключаемся к базе данных
	logger.Info("Подключаемся к БД")
	db, err := database.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
	defer db.Close()

	// Регистрируем gRPC-сервер
	s := grpc.NewServer(grpc.UnaryInterceptor(loggingMiddleware.LoggingError))

	// Регистрируем репозитории
	authRepository := authRepo.New(db, logger)

	// Регистрируем сервисы
	userClient := userService.New(userConn)
	authService := authService.New(authRepository, userClient, logger)

	// Регистрируем эндпоинты
	pbAuth.RegisterAuthServer(s, authEndpoint.New(authService, logger))

	// Запускаем gRPC-сервер
	if err = grpcPkg.ServeGRPC(s, cfg.Services.Auth.GRPC); err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
}
