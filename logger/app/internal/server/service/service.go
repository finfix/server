package service

import (
	"context"
	"time"

	"logger/app/logging"
	"logger/app/logging/enum"
	pb "logger/app/pblogger"
	"pkg/errors"

	"google.golang.org/protobuf/types/known/emptypb"
)

type repository interface {
	AddLog(context.Context, *pb.Log) error
}

type service struct {
	pb.UnsafeLoggerServer
	repository repository
	logger     *logging.Logger
}

func (s *service) AddLog(ctx context.Context, dto *pb.Log) (*emptypb.Empty, error) {
	s.logger.Info("Method AddLog")

	res := &emptypb.Empty{}

	// Валидируем время
	_, err := time.Parse("2006-01-02 15:04:05", dto.Time)
	if err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	// Валидируем уровень лога
	if err := enum.LogLevelValidation(dto.Level); err != nil {
		return res, err
	}

	return res, s.repository.AddLog(ctx, dto)
}

func (s *service) mustEmbedUnimplementedLoggerServer() {}

func New(repository repository, logger *logging.Logger) pb.LoggerServer {
	return &service{
		repository: repository,
		logger:     logger,
	}
}
