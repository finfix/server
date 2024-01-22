package endpoints

import (
	"context"
	"pkg/converter"

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

	out, err := converter.Convert(pb.UpdateCurrenciesRes{}, model.UpdateCurrenciesRes{Rates: rates})
	if err != nil {
		return nil, err
	}

	return &out, nil
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
