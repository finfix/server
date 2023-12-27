package main

import (
	"logger/app/logging"
	loggingMiddleware "logger/app/logging/middleware"
	"logger/app/pblogger"

	"pkg/errors"
	grpcPkg "pkg/grpc"
	"pkg/panicRecover"

	"tgBot/app/internal/client"
	"tgBot/app/internal/config"
	tgBotEndpoint "tgBot/app/internal/services/tgBot/endpoints"
	tgBotService "tgBot/app/internal/services/tgBot/service"
	"tgBot/app/proto/pbTgBot"

	"google.golang.org/grpc"
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
	logging.Init(loggerClient, "tgBot")

	// Регистрируем gRPC-сервер
	s := grpc.NewServer(grpc.UnaryInterceptor(loggingMiddleware.LoggingError))

	// Инициализируем клиента телеграм
	tgBot, tgChat, err := client.Init(cfg.Telegram.Token, cfg.Telegram.ChatID)

	// Регистрируем сервисы
	tgBotService := tgBotService.New(tgBot, tgChat, logger)

	// Регистрируем эндпоинты
	pbTgBot.RegisterTgBotServer(s, tgBotEndpoint.New(tgBotService, logger))

	// Запускаем gRPC-сервер
	if err = grpcPkg.ServeGRPC(s, cfg.Services.TgBot.GRPC); err != nil {
		logger.Fatal(errors.InternalServer.Wrap(err))
	}
}
