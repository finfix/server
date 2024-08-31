package service

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	"server/pkg/errors"
	"server/pkg/log"
	"server/pkg/tgBot"
	settingsModel "server/internal/services/settings/model"
	"server/internal/services/settings/model/applicationType"
	"server/internal/services/settings/network"
	settingsRepository "server/internal/services/settings/repository"
	userModel "server/internal/services/user/model"
	userService "server/internal/services/user/service"
)

var _ SettingsRepository = &settingsRepository.Repository{}
var _ UserService = &userService.Service{}

type SettingsRepository interface {
	UpdateCurrencies(ctx context.Context, rates map[string]decimal.Decimal) error
	GetCurrencies(context.Context) ([]settingsModel.Currency, error)
	GetIcons(context.Context) ([]settingsModel.Icon, error)
	GetVersion(context.Context, applicationType.Type) (settingsModel.Version, error)
}

type UserService interface {
	SendNotification(ctx context.Context, userID uint32, push userModel.Notification) (uint8, error)
	GetUsers(ctx context.Context, filters userModel.GetUsersReq) (users []userModel.User, err error)
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
	userService        UserService
	tgBot              *tgBot.TgBot
	credentials        Credentials
	version            Version
}

// UpdateCurrencies обновляет курсы валют
func (s *Service) UpdateCurrencies(ctx context.Context, req settingsModel.UpdateCurrenciesReq) error {

	err := s.checkAdmin(ctx, req.Necessary.UserID)
	if err != nil {
		return err
	}

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

func (s *Service) SendNotification(ctx context.Context, req settingsModel.SendNotificationReq) (res settingsModel.SendNotificationRes, err error) {

	err = s.checkAdmin(ctx, req.Necessary.UserID)
	if err != nil {
		return res, err
	}

	count, err := s.userService.SendNotification(
		ctx,
		req.UserID,
		req.Notification,
	)
	if err != nil {
		return res, err
	}
	res.NotificationsSent = count
	return res, err
}

func (s *Service) checkAdmin(ctx context.Context, userID uint32) error {
	users, err := s.userService.GetUsers(ctx, userModel.GetUsersReq{ //nolint:exhaustruct
		IDs: []uint32{userID},
	})
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.NotFound.New("User not found")
	}
	user := users[0]
	if !user.IsAdmin {
		return errors.Forbidden.New("Access denied")
	}
	return nil
}

func New(
	settingsRepository SettingsRepository,
	userService UserService,
	tgBot *tgBot.TgBot,
	version Version,
	credentials Credentials,
) *Service {
	return &Service{
		settingsRepository: settingsRepository,
		userService:        userService,
		tgBot:              tgBot,
		credentials:        credentials,
		version:            version,
	}
}
