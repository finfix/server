package endpoints

import (
	"context"
	"tgBot/app/internal/services/tgBot/endpoints/converter"
	"tgBot/app/internal/services/tgBot/model"
	tgBotService "tgBot/app/internal/services/tgBot/service"
	pb "tgBot/app/proto/pbTgBot"

	"logger/app/logging"

	"google.golang.org/protobuf/types/known/emptypb"
)

var _ TgBotService = &tgBotService.Service{}

type TgBotService interface {
	SendMessage(context.Context, model.SendMessageReq) error
}

func (s *Endpoint) SendMessage(ctx context.Context, in *pb.SendMessageReq) (*emptypb.Empty, error) {

	if err := s.service.SendMessage(ctx, converter.PbSendMessageReq{in}.ConvertToStruct()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type Endpoint struct {
	pb.UnsafeTgBotServer
	service TgBotService
	logger  *logging.Logger
}

func New(service TgBotService, logger *logging.Logger) pb.TgBotServer {
	return &Endpoint{
		service: service,
		logger:  logger,
	}
}
