package main

import (
	grpcPkg "pkg/grpc"

	"google.golang.org/grpc"

	"logger/app/internal/config"
	"logger/app/internal/server/repository"
	"logger/app/internal/server/service"
	pb "logger/app/pblogger"
	"pkg/errors"
	"pkg/panicRecover"

	"logger/app/logging"
	"logger/app/logging/middleware"
	"pkg/database"
)

func main() {

	// Перехватываем панику
	defer panicRecover.PanicRecover(func(err error) {
		logging.GetLogger().Panic(err)
	})

	logger := logging.GetLogger()

	cfg := config.GetConfig()

	// Configure logger
	logging.Init(nil, "logger")

	logger.Info("Connect to DB")
	dbx, err := database.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		logger.Fatal(err)
	}

	// Create gRPC server
	s := grpc.NewServer(grpc.UnaryInterceptor(middleware.LoggingError))

	pb.RegisterLoggerServer(s, service.New(repository.New(dbx, logger), logger))

	// Запускаем gRPC-сервер
	if err = grpcPkg.ServeGRPC(s, cfg.Services.Logger.GRPC); err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
}
