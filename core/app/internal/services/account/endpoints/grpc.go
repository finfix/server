package endpoints

import (
	"context"

	"logger/app/logging"

	"core/app/internal/services/account/endpoints/converter"
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

	id, err := s.service.Create(context.Background(), converter.PbCreateReq{CreateReq: in}.ConvertToStruct())
	if err != nil {
		return nil, err
	}

	res := model.CreateRes{ID: id}
	return converter.CreateRes{res}.ConvertToProto(), nil
}

func (s *Endpoint) Get(ctx context.Context, in *pb.GetReq) (*pb.GetRes, error) {
	s.logger.Info("Method Get")

	accounts, err := s.service.Get(ctx, converter.PbGetReq{GetReq: in}.ConvertToStruct())
	if err != nil {
		return nil, err
	}

	res := model.GetRes{Accounts: accounts}
	return converter.GetRes{res}.ConvertToProto(), nil
}

func (s *Endpoint) Update(_ context.Context, in *pb.UpdateReq) (*emptypb.Empty, error) {
	s.logger.Info("Method Update")

	if err := s.service.Update(context.Background(), converter.PbUpdateReq{UpdateReq: in}.ConvertToStruct()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Endpoint) Delete(_ context.Context, in *pb.DeleteReq) (*emptypb.Empty, error) {
	s.logger.Info("Method Delete")

	if err := s.service.Delete(context.Background(), converter.PbDeleteReq{DeleteReq: in}.ConvertToStruct()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Endpoint) Switch(_ context.Context, in *pb.SwitchReq) (*emptypb.Empty, error) {
	s.logger.Info("Method Switch")

	if err := s.service.Switch(context.Background(), converter.PbSwitchReq{SwitchReq: in}.ConvertToStruct()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Endpoint) GetAccountGroups(ctx context.Context, req *pb.GetAccountGroupsReq) (*pb.GetAccountGroupsRes, error) {
	s.logger.Info("Method GetAccountGroups")

	accountGroups, err := s.service.GetAccountGroups(ctx, converter.PbGetAccountGroupsReq{GetAccountGroupsReq: req}.ConvertToStruct())
	if err != nil {
		return nil, err
	}

	res := model.GetAccountGroupsRes{AccountGroups: accountGroups}
	return converter.GetAccountGroupsRes{res}.ConvertToProto(), nil
}

func New(service AccountService, logger *logging.Logger) pb.AccountServer {
	return &Endpoint{
		service: service,
		logger:  logger,
	}
}
