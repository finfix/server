package endpoints

import (
	"context"
	"pkg/converter"

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

func (s *Endpoint) Get(ctx context.Context, in *pb.GetReq) (*pb.GetRes, error) {
	s.logger.Info("Method Get")

	req, err := converter.Convert(model.GetReq{}, in)
	if err != nil {
		return nil, err
	}

	transactions, err := s.service.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	out, err := converter.Convert(pb.GetRes{}, model.GetRes{Transactions: transactions})
	if err != nil {
		return nil, err
	}

	return &out, nil
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
