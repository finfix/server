package service

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	"server/app/pkg/errors"
	"server/app/pkg/log"
	"server/app/pkg/tgBot"
	settingsModel "server/app/services/settings/model"
	"server/app/services/settings/model/applicationType"
	"server/app/services/settings/network"
	settingsRepository "server/app/services/settings/repository"
)

var _ SettingsRepository = &settingsRepository.Repository{}

type SettingsRepository interface {
	UpdateCurrencies(ctx context.Context, rates map[string]decimal.Decimal) error
	GetCurrencies(context.Context) ([]settingsModel.Currency, error)
	GetIcons(context.Context) ([]settingsModel.Icon, error)
	GetVersion(context.Context, applicationType.Type) (settingsModel.Version, error)
}

type Credentials struct {
	CurrencyProviderAPIKey string
}

type Version struct {
	Version string
	Build   string
}

type Service struct {
	settingsRepository SettingsRepository
	tgBot              *tgBot.TgBot
	credentials        Credentials
	version            Version
}

// UpdateCurrencies обновляет курсы валют
func (s *Service) UpdateCurrencies(ctx context.Context) error {

	const updateCurrenciesTemplate = "*📈 Курс валют успешно обновлен*\n\nUSD: %v₽\nBTC: %v$"

	var tgMessage tgBot.SendMessageReq

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

func (s *Service) GetVersion(ctx context.Context, appType applicationType.Type) (version settingsModel.Version, err error) {
	switch appType {
	case applicationType.Server:
		return settingsModel.Version{
			Version: s.version.Version,
			Build:   s.version.Build,
		}, nil
	case applicationType.IOs:
		return s.settingsRepository.GetVersion(ctx, appType)
	case applicationType.Android, applicationType.Web:
		return version, errors.NotFound.New("Такое приложение еще не реализовано")
	default:
		return version, errors.BadRequest.New("Неверный тип приложения")
	}
}

func New(rep SettingsRepository, tgBot *tgBot.TgBot, version Version, credentials Credentials) *Service {
	return &Service{
		settingsRepository: rep,
		tgBot:              tgBot,
		credentials:        credentials,
		version:            version,
	}
}
