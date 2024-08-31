package service

import (
	"context"
	"strings"

	"server/internal/services/tgBot/model"
	"server/pkg/errors"
	"server/pkg/log"
)

// SendMessage отправляет сообщение пользователю в телеграм
func (s *TgBotService) SendMessage(ctx context.Context, req model.SendMessageReq) error {

	if !s.isOn {
		log.Warning(ctx, "Вызвана функция SendMessage. Пуши выключены")
		return nil
	}

	req.Message = strings.ReplaceAll(req.Message, ".", "\\.")

	if _, err := s.Bot.Send(s.Chat, req.Message); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}
