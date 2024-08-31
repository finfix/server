package service

import (
	"context"

	"github.com/shopspring/decimal"

	settingsModel "server/internal/services/settings/model"
	"server/internal/services/settings/model/applicationType"
	settingsRepository "server/internal/services/settings/repository"
	tgBotModel "server/internal/services/tgBot/model"
	userModel "server/internal/services/user/model"
	userService "server/internal/services/user/service"
)

var _ SettingsRepository = &settingsRepository.SettingsRepository{}
var _ UserService = &userService.UserService{}

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

type TgBotService interface {
	SendMessage(context.Context, tgBotModel.SendMessageReq) error
}

type Credentials struct {
	CurrencyProviderAPIKey string
}

type Version struct {
	Version string
	Build   string
}

type SettingsService struct {
	settingsRepository SettingsRepository
	userService        UserService
	tgBot              TgBotService
	credentials        Credentials
	version            Version
}

func NewSettingsService(
	settingsRepository SettingsRepository,
	userService UserService,
	tgBot TgBotService,
	version Version,
	credentials Credentials,
) *SettingsService {
	return &SettingsService{
		settingsRepository: settingsRepository,
		userService:        userService,
		tgBot:              tgBot,
		credentials:        credentials,
		version:            version,
	}
}
