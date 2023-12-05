package endpoints

import (
	"context"

	"logger/app/logging"

	"core/app/internal/services/user/endpoints/converter"
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

	return converter.GetCurrenciesRes{&model.GetCurrenciesRes{Currencies: currencies}}.ConvertToProto(), nil
}

func (s *Endpoint) Create(ctx context.Context, in *pb.CreateReq) (*pb.CreateRes, error) {
	s.logger.Info("Method Create")

	id, err := s.service.Create(ctx, converter.PbCreateReq{in}.ConvertToStruct())
	if err != nil {
		return nil, err
	}

	return converter.CreateRes{&model.CreateRes{id}}.ConvertToProto(), nil
}

func (s *Endpoint) Get(ctx context.Context, in *pb.GetReq) (*pb.GetRes, error) {
	s.logger.Info("Method Get")

	user, err := s.service.Get(ctx, converter.PbGetReq{GetReq: in}.ConvertToStruct())
	if err != nil {
		return nil, err
	}

	return converter.GetRes{&model.GetRes{Users: user}}.ConvertToProto(), nil
}

func New(service UserService, logger *logging.Logger) pb.UserServer {
	return &Endpoint{
		service: service,
		logger:  logger,
	}
}
