package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/pkg/log"
	"server/app/pkg/passwordManager"
	"server/app/pkg/pushNotificator"
	"server/app/services/generalRepository"
	userModel "server/app/services/user/model"
	"server/app/services/user/model/OS"
	userRepository "server/app/services/user/repository"
	userRepoModel "server/app/services/user/repository/model"
)

var _ UserRepository = &userRepository.Repository{}
var _ GeneralRepository = &generalRepository.Repository{}

type GeneralRepository interface {
	WithinTransaction(ctx context.Context, callback func(context.Context) error) error
}

type UserRepository interface {
	CreateUser(context.Context, userModel.CreateReq) (uint32, error)
	GetUsers(context.Context, userModel.GetReq) ([]userModel.User, error)
	UpdateUser(context.Context, userRepoModel.UpdateUserReq) error

	LinkUserToAccountGroup(context.Context, uint32, uint32) error

	GetDevices(context.Context, userRepoModel.GetDevicesReq) ([]userModel.Device, error)
	UpdateDevice(context.Context, userRepoModel.UpdateDeviceReq) error
}

type Service struct {
	userRepository    UserRepository
	generalRepository GeneralRepository
	pushNotificator   *pushNotificator.PushNotificator
	generalSalt       []byte
}

// CreateUser создает нового пользователя
func (s *Service) CreateUser(ctx context.Context, user userModel.CreateReq) (id uint32, err error) {
	return s.userRepository.CreateUser(ctx, user)
}

// GetUsers возвращает всех юзеров по фильтрам
func (s *Service) GetUsers(ctx context.Context, filters userModel.GetReq) (users []userModel.User, err error) {
	return s.userRepository.GetUsers(ctx, filters)
}

// UpdateUser обновляет настройки пользователя
func (s *Service) UpdateUser(ctx context.Context, req userModel.UpdateUserReq) error {

	return s.generalRepository.WithinTransaction(ctx, func(ctx context.Context) error {

		// Если обновляется токен уведомлений, обновляем его в таблице девайсов
		if req.NotificationToken != nil {
			if err := s.userRepository.UpdateDevice(ctx, userRepoModel.UpdateDeviceReq{
				UserID:            req.Necessary.UserID,
				DeviceID:          req.Necessary.DeviceID,
				RefreshToken:      nil,
				NotificationToken: req.NotificationToken,
				ApplicationInformation: userRepoModel.UpdateApplicationInformationReq{
					BundleID: nil,
					Version:  nil,
					Build:    nil,
				},
				DeviceInformation: userRepoModel.UpdateDeviceInformationReq{
					VersionOS: nil,
					IPAddress: nil,
					UserAgent: nil,
				},
			}); err != nil {
				return err
			}
		}

		repoReq := req.ConvertToRepoModel()

		// Если обновляется пароль
		if req.Password != nil {

			if req.OldPassword != nil {
				return errors.BadRequest.New("OldPassword must be filled")
			}

			// Получаем актуальный пароль пользователя
			users, err := s.userRepository.GetUsers(ctx, userModel.GetReq{ //nolint:exhaustruct
				IDs: []uint32{req.Necessary.UserID},
			})
			if err != nil {
				return err
			}
			if len(users) == 0 {
				return errors.NotFound.New("User not found")
			}
			user := users[0]

			// Сравниваем пришедший пароль и хэш пароля из базы данных
			if err = passwordManager.CompareHashAndPassword(user.PasswordHash, []byte(*req.OldPassword), user.PasswordSalt, s.generalSalt); err != nil {
				return err
			}

			// Получаем хэш и соль нового пароля
			passwordHash, passwordSalt, err := passwordManager.CreateNewPassword([]byte(*req.Password), s.generalSalt)
			if err != nil {
				return err
			}

			repoReq.PasswordHash = &passwordHash
			repoReq.PasswordSalt = &passwordSalt
		}

		if err := s.userRepository.UpdateUser(ctx, repoReq); err != nil {
			return err
		}

		return nil
	})
}

// SendNotification отправляет пуш на все устройства пользователя
func (s *Service) SendNotification(ctx context.Context, userID uint32, push userModel.Notification) (count uint8, err error) {

	// Получаем все девайсы пользователя
	devices, err := s.userRepository.GetDevices(ctx, userRepoModel.GetDevicesReq{ //nolint:exhaustruct
		UserIDs: []uint32{userID},
	})
	if err != nil {
		return count, err
	}

	// Проходимся по каждому девайсу
	for _, device := range devices {

		if device.NotificationToken == nil {
			continue
		}

		// Смотрим на операционную систему и отправляем уведомление
		switch device.DeviceInformation.NameOS {
		case OS.IOS, OS.IPadOS, OS.OSX, OS.WatchOS:
			_, err = s.pushNotificator.Push(ctx, pushNotificator.PushReq{
				Notification: pushNotificator.NotificationSettings{
					Title:    &push.Title,
					Message:  &push.Message,
					Subtitle: &push.Subtitle,
					Badge:    &push.BadgeCount,
				},
				NotificationToken: *device.NotificationToken,
				BundleID:          device.ApplicationInformation.BundleID,
			})

		case OS.Android:
			break
		}
		if err != nil {
			log.Error(ctx, err)
		}
		count++
	}

	return count, nil
}

func New(
	userRepository UserRepository,
	generalRepository GeneralRepository,
	pushNotificator *pushNotificator.PushNotificator,
	generalSalt []byte,
) *Service {
	return &Service{
		userRepository:    userRepository,
		generalRepository: generalRepository,
		pushNotificator:   pushNotificator,
		generalSalt:       generalSalt,
	}
}
