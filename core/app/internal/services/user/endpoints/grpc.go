package endpoints

import (
	"context"
	"pkg/converter"

	"logger/app/logging"

	"core/app/internal/services/user/model"
	userService "core/app/internal/services/user/service"
	pb "core/app/proto/pbUser"

	"google.golang.org/protobuf/types/known/emptypb"
)

var _ UserService = &userService.Service{}

type UserService interface {
	Create(context.Context, model.CreateReq) (id uint32, err error)
	Get(context.Context, model.GetReq) ([]model.User, error)
	GetCurrencies(context.Context) ([]model.Currency, error)
}

type Endpoint struct {
	pb.UnsafeUserServer
	service UserService
	logger  *logging.Logger
}

func (s *Endpoint) GetCurrencies(ctx context.Context, empty *emptypb.Empty) (*pb.GetCurrenciesRes, error) {
	s.logger.Info("Method GetCurrencies")

	currencies, err := s.service.GetCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	out, err := converter.Convert(pb.GetCurrenciesRes{}, model.GetCurrenciesRes{Currencies: currencies})
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (s *Endpoint) Create(ctx context.Context, in *pb.CreateReq) (*pb.CreateRes, error) {
	s.logger.Info("Method Create")

	req, err := converter.Convert(model.CreateReq{}, in)
	if err != nil {
		return nil, err
	}

	id, err := s.service.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	out, err := converter.Convert(pb.CreateRes{}, model.CreateRes{ID: id})
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (s *Endpoint) Get(ctx context.Context, in *pb.GetReq) (*pb.GetRes, error) {
	s.logger.Info("Method Get")

	req, err := converter.Convert(model.GetReq{}, in)
	if err != nil {
		return nil, err
	}

	users, err := s.service.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	out, err := converter.Convert(pb.GetRes{}, model.GetRes{Users: users})
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func New(service UserService, logger *logging.Logger) pb.UserServer {
	return &Endpoint{
		service: service,
		logger:  logger,
	}
}
