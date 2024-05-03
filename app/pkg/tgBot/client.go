package tgBot

import (
	"context"
	"strings"

	"gopkg.in/telebot.v3"

	"server/app/pkg/errors"
	"server/app/pkg/logging"
)

type TgBot struct {
	Bot    *telebot.Bot
	Chat   *telebot.Chat
	logger *logging.Logger
	isOn   bool
}

func NewTgBot(
	token string,
	chatID int64,
	isOn bool,
	logger *logging.Logger,
) (*TgBot, error) {

	if !isOn {
		logger.Debug(context.Background(), "Telegram bot is off")
		return &TgBot{
			Bot:    nil,
			Chat:   nil,
			isOn:   isOn,
			logger: logger,
		}, nil
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token: token,
	})
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	chat, err := bot.ChatByID(chatID)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	return &TgBot{
		Bot:    bot,
		Chat:   chat,
		isOn:   isOn,
		logger: logger,
	}, nil
}

// SendMessage отправляет сообщение пользователю в телеграм
func (s *TgBot) SendMessage(_ context.Context, req SendMessageReq) error {

	if !s.isOn {
		return nil
	}

	opts := &telebot.SendOptions{
		ParseMode: telebot.ModeMarkdownV2,
	}

	req.Message = strings.ReplaceAll(req.Message, ".", "\\.")

	if _, err := s.Bot.Send(s.Chat, req.Message, opts); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}
