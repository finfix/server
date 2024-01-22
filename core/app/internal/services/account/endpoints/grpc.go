package endpoints

import (
	"context"
	converter "pkg/converter"

	"logger/app/logging"

	"core/app/internal/services/account/model"
	accountService "core/app/internal/services/account/service"
	pb "core/app/proto/pbAccount"

	"google.golang.org/protobuf/types/known/emptypb"
)

var _ AccountService = &accountService.Service{}

type AccountService interface {
	Create(context.Context, model.CreateReq) (id uint32, err error)
	Get(context.Context, model.GetReq) ([]model.Account, error)
	Update(context.Context, model.UpdateReq) error
	Delete(context.Context, model.DeleteReq) error
	Switch(context.Context, model.SwitchReq) error
	GetAccountGroups(context.Context, model.GetAccountGroupsReq) ([]model.AccountGroup, error)
}

type Endpoint struct {
	pb.UnsafeAccountServer
	service AccountService
	logger  *logging.Logger
}

func (s *Endpoint) Create(_ context.Context, in *pb.CreateReq) (*pb.CreateRes, error) {
	s.logger.Info("Method Create")

	req, err := converter.Convert(model.CreateReq{}, in)
	if err != nil {
		return nil, err
	}

	id, err := s.service.Create(context.Background(), req)
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

	accounts, err := s.service.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	out, err := converter.Convert(pb.GetRes{}, model.GetRes{Accounts: accounts})
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (s *Endpoint) Update(_ context.Context, in *pb.UpdateReq) (*emptypb.Empty, error) {
	s.logger.Info("Method Update")

	req, err := converter.Convert(model.UpdateReq{}, in)
	if err != nil {
		return nil, err
	}

	if err := s.service.Update(context.Background(), req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Endpoint) Delete(_ context.Context, in *pb.DeleteReq) (*emptypb.Empty, error) {
	s.logger.Info("Method Delete")

	req, err := converter.Convert(model.DeleteReq{}, in)
	if err != nil {
		return nil, err
	}

	if err := s.service.Delete(context.Background(), req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Endpoint) Switch(_ context.Context, in *pb.SwitchReq) (*emptypb.Empty, error) {
	s.logger.Info("Method Switch")

	req, err := converter.Convert(model.SwitchReq{}, in)
	if err != nil {
		return nil, err
	}

	if err := s.service.Switch(context.Background(), req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Endpoint) GetAccountGroups(ctx context.Context, in *pb.GetAccountGroupsReq) (*pb.GetAccountGroupsRes, error) {
	s.logger.Info("Method GetAccountGroups")

	req, err := converter.Convert(model.GetAccountGroupsReq{}, in)
	if err != nil {
		return nil, err
	}

	accountGroups, err := s.service.GetAccountGroups(ctx, req)
	if err != nil {
		return nil, err
	}

	out, err := converter.Convert(pb.GetAccountGroupsRes{}, model.GetAccountGroupsRes{AccountGroups: accountGroups})
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func New(service AccountService, logger *logging.Logger) pb.AccountServer {
	return &Endpoint{
		service: service,
		logger:  logger,
	}
}
