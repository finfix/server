package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"testing"

	"server/app/services/account/model"
	"server/app/services/account/service/mocks"
	"server/app/services/generalRepository/checker"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/testingFunc"
)

func TestService_Update(t *testing.T) {

	logging.Off()

	type interfaces struct {
		accountRepository     *mocks.AccountRepository
		transactionRepository *mocks.TransactionRepository
		generalRepository     *mocks.GeneralRepository
		userRepository        *mocks.UserRepository
		permissionsService    *mocks.PermissionsService
	}

	var (
		ctx = context.Background()
	)

	tests := []struct {
		name          string
		accountFields model.UpdateReq
		wantError     error
		mockActions   func(interfaces)
	}{
		{
			name: "1. Пользователь имеет доступ к счету и редактирование счета прошло успешно",
			accountFields: model.UpdateReq{
				ID:     1,
				UserID: 1,
			},
			wantError: nil,
			mockActions: func(inf interfaces) {
				// Пользователь имеет доступ к счету
				inf.generalRepository.On("CheckAccess", ctx, checker.Accounts, uint32(1), []uint32{1}).Return(nil)
				inf.generalRepository.On("WithinTransaction", ctx, mock.Anything).Return(nil)
			},
		},
		{
			name: "2. Пользователь имеет доступ к счету, но произошла ошибка при редактировании счета",
			accountFields: model.UpdateReq{
				ID:     1,
				UserID: 1,
			},
			wantError: errors.InternalServer.New("Ошибка при редактировании счета"),
			mockActions: func(inf interfaces) {
				// Пользователь имеет доступ к счету
				inf.generalRepository.On("CheckAccess", ctx, checker.Accounts, uint32(1), []uint32{1}).Return(nil)
				inf.generalRepository.On("WithinTransaction", ctx, mock.Anything).Return(errors.InternalServer.New("Ошибка при редактировании счета"))
			},
		},
		{
			name: "3. Пользователь не имеет доступ к счету",
			accountFields: model.UpdateReq{
				ID:     1,
				UserID: 1,
			},
			wantError: errors.InternalServer.New("Пользователь не имеет доступ к счету"),
			mockActions: func(inf interfaces) {
				// Пользователь не имеет доступ к счету
				inf.generalRepository.On("CheckAccess", ctx, checker.Accounts, uint32(1), []uint32{1}).Return(errors.InternalServer.New("Пользователь не имеет доступ к счету"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountRepository := mocks.NewAccountRepository(t)
			generalRepository := mocks.NewGeneralRepository(t)
			transactionRepository := mocks.NewTransactionRepository(t)
			userRepository := mocks.NewUserRepository(t)
			permissionService := mocks.NewPermissionsService(t)

			s := &Service{
				accountRepository:  accountRepository,
				general:            generalRepository,
				transaction:        transactionRepository,
				user:               userRepository,
				permissionsService: permissionService,
				logger:             logging.GetLogger(),
			}
			interfaces := interfaces{
				accountRepository:     accountRepository,
				transactionRepository: transactionRepository,
				generalRepository:     generalRepository,
				userRepository:        userRepository,
				permissionsService:    permissionService,
			}

			tt.mockActions(interfaces)

			err := s.Update(context.Background(), tt.accountFields)
			testingFunc.CheckError(t, err, tt.wantError)
		})
	}
}
