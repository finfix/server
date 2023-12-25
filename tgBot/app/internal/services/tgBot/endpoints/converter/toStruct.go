package converter

import (
	"tgBot/app/internal/services/tgBot/model"
	"tgBot/app/proto/pbTgBot"
)

type PbSendMessageReq struct {
	*pbTgBot.SendMessageReq
}

func (pb PbSendMessageReq) ConvertToStruct() model.SendMessageReq {
	var p model.SendMessageReq
	p.Message = pb.Message
	return p
}
