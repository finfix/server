package service

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"

	"server/app/pkg/logging"
	settingsModel "server/app/services/settings/model"
	"server/app/services/settings/network"
	settingsRepository "server/app/services/settings/repository"
	tgBotModel "server/app/services/tgBot/model"
	tgBotService "server/app/services/tgBot/service"
)

var _ SettingsRepository = &settingsRepository.Repository{}
var _ TgBotService = &tgBotService.Service{}

type SettingsRepository interface {
	UpdateCurrencies(ctx context.Context, rates map[string]decimal.Decimal) error
	GetCurrencies(context.Context) ([]settingsModel.Currency, error)
	GetIcons(context.Context) ([]settingsModel.Icon, error)
}

type TgBotService interface {
	SendMessage(context.Context, tgBotModel.SendMessageReq) error
}

type Service struct {
	settingsRepository SettingsRepository
	tgBotService       TgBotService
	logger             *logging.Logger
	version            string
	build              string
}

// UpdateCurrencies обновляет курсы валют
func (s *Service) UpdateCurrencies(ctx context.Context) error {

	const updateCurrenciesTemplate = "*📈 Курс валют успешно обновлен*\n\nUSD: %v₽\nBTC: %v$"

	var tgMessage tgBotModel.SendMessageReq

	defer func() {
		err := s.tgBotService.SendMessage(ctx, tgMessage)
		if err != nil {
			s.logger.Error(ctx, err)
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
	if err = s.settingsRepository.UpdateCurrencies(ctx, rates); err != nil {
		tgMessage.Message += fmt.Sprintf("Не смогли обновить курсы валют в базе данных\n\n%v", err.Error())
		return err
	}

	tgMessage.Message = fmt.Sprintf(
		updateCurrenciesTemplate,
		getRate(rates, "USD", "RUB").Round(2), //nolint:gomnd
		getRate(rates, "BTC", "USD").Round(0),
	)

	return nil
}

func getRate(rates map[string]decimal.Decimal, currency, currencyRelate string) decimal.Decimal {
	currencyRate := rates[currency]
	currencyRelateRate := rates[currencyRelate]
	if currencyRate.GreaterThan(currencyRelateRate) {
		return currencyRate.Div(currencyRelateRate)

	}
	return currencyRelateRate.Div(currencyRate)
}

func (s *Service) GetCurrencies(ctx context.Context) ([]settingsModel.Currency, error) {
	return s.settingsRepository.GetCurrencies(ctx)
}

func (s *Service) GetIcons(ctx context.Context) ([]settingsModel.Icon, error) {
	return s.settingsRepository.GetIcons(ctx)
}

func (s *Service) GetVersion() settingsModel.Version {
	return settingsModel.Version{
		Version: s.version,
		Build:   s.build,
	}
}

func New(rep SettingsRepository, tgBotService TgBotService, logger *logging.Logger, version, build string) *Service {
	return &Service{
		settingsRepository: rep,
		tgBotService:       tgBotService,
		logger:             logger,
		version:            version,
		build:              build,
	}
}
