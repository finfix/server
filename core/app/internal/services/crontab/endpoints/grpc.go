package endpoints

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"logger/app/logging"

	model "core/app/internal/services/crontab/model"
	pb "core/app/proto/pbCrontab"
)

type CrontabService interface {
	UpdateCurrencies(context.Context) (map[string]float64, error)
}

func (s *Endpoint) UpdateCurrencies(ctx context.Context, _ *emptypb.Empty) (*pb.UpdateCurrenciesRes, error) {
	s.logger.Info("Method UpdateCurrencies")

	rates, err := s.service.UpdateCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	res := model.UpdateCurrenciesRes{Rates: rates}
	return res.ConvertToProto(), nil
}

type Endpoint struct {
	pb.UnsafeCrontabServer
	service CrontabService
	logger  *logging.Logger
}

func New(service CrontabService, logger *logging.Logger) pb.CrontabServer {
	return &Endpoint{
		service: service,
		logger:  logger,
	}
}
