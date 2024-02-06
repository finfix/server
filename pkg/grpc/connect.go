package grpc

import (
	"net"
	"pkg/errors"

	"google.golang.org/grpc"
)

func ServeGRPC(s *grpc.Server, port string) error {

	// Проверяем наличие порта для запуска gRPC-сервера
	if port == "" {
		return errors.InternalServer.New("Переменная окружения с портом сервиса не задана")
	}

	// Начинаем слушать порт
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	// Запускаем gRPC-сервер
	if err := s.Serve(lis); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}
