package service

import (
	"context"
	"fmt"

	"server/app/services/admin/network"
	tgBotModel "server/app/services/tgBot/model"
	"server/pkg/logging"
)

type Repository interface {
	UpdCurrencies(ctx context.Context, rates map[string]float64) error
}

type TgBotService interface {
	SendMessage(context.Context, tgBotModel.SendMessageReq) error
}

// UpdateCurrencies обновляет курсы валют
func (s *Service) UpdateCurrencies(ctx context.Context) error {

	const updateCurrenciesTemplate = "*📈 Курс валют успешно обновлен*\n\nUSD: %.2f₽\nBTC: %.0f$"

	var tgMessage tgBotModel.SendMessageReq

	defer func() {
		err := s.tgBotService.SendMessage(ctx, tgMessage)
		if err != nil {
			s.logger.Error(err)
		}
	}()

	// Получаем курсы валют от провайдера данных
	rates, err := network.GetCurrencyRates(ctx)
	if err != nil {
		tgMessage.Message += fmt.Sprintf("Не смогли получить курсы валют от провайдера\n\n%v", err.Error())
		return err
	}
	tgMessage.Message += "Успешно получили курсы валют от провайдера\n"

	// Обновляем курсы валют в БД
	if err = s.repository.UpdCurrencies(ctx, rates); err != nil {
		tgMessage.Message += fmt.Sprintf("Не смогли обновить курсы валют в базе данных\n\n%v", err.Error())
		return err
	}

	tgMessage.Message = fmt.Sprintf(
		updateCurrenciesTemplate,
		getRate(rates, "USD", "RUB"),
		getRate(rates, "BTC", "USD"),
	)

	return nil
}

func getRate(rates map[string]float64, currency, currencyRelate string) float64 {
	currencyRate := rates[currency]
	currencyRelateRate := rates[currencyRelate]
	if currencyRate > currencyRelateRate {
		return currencyRate / currencyRelateRate

	}
	return currencyRelateRate / currencyRate
}

type Service struct {
	repository   Repository
	tgBotService TgBotService
	logger       *logging.Logger
}

func New(rep Repository, tgBotService TgBotService, logger *logging.Logger) *Service {
	return &Service{
		repository:   rep,
		tgBotService: tgBotService,
		logger:       logger,
	}
}
