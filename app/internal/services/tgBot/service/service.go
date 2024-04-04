package service

import (
	"context"
	"strings"

	"gopkg.in/telebot.v3"

	"server/app/internal/services/tgBot/model"
	"server/pkg/errors"
	"server/pkg/logging"
)

// SendMessage отправляет сообщение пользователю в телеграм
func (s *Service) SendMessage(_ context.Context, req model.SendMessageReq) error {

	opts := &telebot.SendOptions{
		ParseMode: telebot.ModeMarkdownV2,
	}

	req.Message = strings.ReplaceAll(req.Message, ".", "\\.")

	if _, err := s.bot.Send(s.chat, req.Message, opts); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}

type Service struct {
	bot    *telebot.Bot
	chat   *telebot.Chat
	logger *logging.Logger
}

func New(
	tgBot *telebot.Bot,
	user *telebot.Chat,
	logger *logging.Logger,
) *Service {
	return &Service{
		bot:    tgBot,
		chat:   user,
		logger: logger,
	}
}
