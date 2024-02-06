package service

import (
	"context"

	"logger/app/logging"

	"tgBot/app/internal/services/tgBot/model"

	"gopkg.in/telebot.v3"
)

// SendMessage отправляет сообщение пользователю в телеграм
func (s *Service) SendMessage(ctx context.Context, req model.SendMessageReq) error {

	if _, err := s.bot.Send(s.chat, req.Message); err != nil {
		return err
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
