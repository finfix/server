package tgBot

import (
	"gopkg.in/telebot.v3"

	"server/pkg/errors"
)

func Init(token string, chatID int64) (*telebot.Bot, *telebot.Chat, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token: token,
	})
	if err != nil {
		return nil, nil, errors.InternalServer.Wrap(err)
	}

	chat, err := bot.ChatByID(chatID)
	if err != nil {
		return nil, nil, errors.InternalServer.Wrap(err)
	}

	return bot, chat, nil
}
