package service

import (
	"context"
	"fmt"

	"pkg/log"

	settingsModel "server/internal/services/settings/model"
	"server/internal/services/settings/network"
	"server/internal/services/settings/service/utils"
	"server/internal/services/tgBot/model"
)

// UpdateCurrencies обновляет курсы валют
func (s *SettingsService) UpdateCurrencies(ctx context.Context, req settingsModel.UpdateCurrenciesReq) error {

	// Проверяем, что пользователь администратор
	err := s.checkAdmin(ctx, req.Necessary.UserID)
	if err != nil {
		return err
	}

	const updateCurrenciesTemplate = "*📈 Курс валют успешно обновлен*\n\nUSD: %v₽\nBTC: %v$"

	var tgMessage model.SendMessageReq

	defer func() {
		err := s.tgBot.SendMessage(ctx, tgMessage)
		if err != nil {
			log.Error(ctx, err)
		}
	}()

	// Получаем курсы валют от провайдера данных
	rates, err := network.GetCurrencyRates(ctx, s.credentials.CurrencyProviderAPIKey)
	if err != nil {
		tgMessage.Message += fmt.Sprintf("Не смогли получить курсы валют от провайдера\n\n%v", err.Error())
		return err
	}
	tgMessage.Message += "Успешно получили курсы валют от провайдера\n"

	// Обновляем курсы валют в БД
	if err = s.settingsRepository.UpdateCurrencies(ctx, rates); err != nil {
		tgMessage.Message += fmt.Sprintf("Не смогли обновить курсы валют в базе данных\n\n%v", err.Error())
		return err
	}

	tgMessage.Message = fmt.Sprintf(
		updateCurrenciesTemplate,
		utils.GetRate(rates, "USD", "RUB").Round(2), //nolint:gomnd
		utils.GetRate(rates, "BTC", "USD").Round(0),
	)

	return nil
}
