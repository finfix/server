package endpoints

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"logger/app/logging"

	"core/app/internal/services/transaction/model"
	pb "core/app/proto/pbTransaction"
)

type TransactionService interface {
	Create(context.Context, model.CreateReq) (id uint32, err error)
	Get(context.Context, model.GetReq) ([]model.Transaction, error)
	Update(context.Context, model.UpdateReq) error
	Delete(context.Context, model.DeleteReq) error
}

func (s *Endpoint) Create(_ context.Context, in *pb.CreateReq) (*pb.CreateRes, error) {
	s.logger.Info("Method Create")

	req := model.PbCreateReq{CreateReq: in}
	id, err := s.service.Create(context.Background(), *req.ConvertToStruct())
	if err != nil {
		return nil, err
	}

	res := model.CreateRes{ID: id}
	return res.ConvertToProto(), nil
}

func (s *Endpoint) Update(_ context.Context, in *pb.UpdateReq) (*emptypb.Empty, error) {
	s.logger.Info("Method Update")

	req := model.PbUpdateReq{UpdateReq: in}
	if err := s.service.Update(context.Background(), *req.ConvertToStruct()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Endpoint) Delete(_ context.Context, in *pb.DeleteReq) (*emptypb.Empty, error) {
	s.logger.Info("Method Delete")

	req := model.PbDeleteReq{DeleteReq: in}
	if err := s.service.Delete(context.Background(), *req.ConvertToStruct()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Endpoint) Get(ctx context.Context, in *pb.GetReq) (*pb.GetRes, error) {
	s.logger.Info("Method Get")

	req := model.PbGetReq{GetReq: in}
	transactions, err := s.service.Get(ctx, *req.ConvertToStruct())
	if err != nil {
		return nil, err
	}

	res := model.GetRes{Transactions: transactions}
	return res.ConvertToProto(), nil
}

type Endpoint struct {
	pb.UnsafeTransactionServer
	service TransactionService
	logger  *logging.Logger
}

func New(service TransactionService, logger *logging.Logger) pb.TransactionServer {
	return &Endpoint{
		service: service,
		logger:  logger,
	}
}
