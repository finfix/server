package service

import (
	"context"
	"testing"

	"server/app/services/account/model"
	"server/app/services/account/service/mocks"
	"server/app/services/permissions"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/pointer"
	"server/pkg/testingFunc"
)

func TestService_update(t *testing.T) {

	logging.Off()

	type interfaces struct {
		accountRepository     *mocks.AccountRepository
		transactionRepository *mocks.TransactionRepository
		generalRepository     *mocks.GeneralRepository
		userRepository        *mocks.UserRepository
		permissionsService    *mocks.PermissionsService
		accountService        *mocks.AccountService
	}

	var (
		ctx       = context.Background()
		updateReq = model.UpdateReq{
			ID:              1,
			UserID:          1,
			Remainder:       pointer.Pointer(1.1),
			ParentAccountID: pointer.Pointer(uint32(2)),
		}
		getAccountsReq = model.GetReq{
			IDs: []uint32{1},
		}
		getAccountRes = []model.Account{
			{
				Name: "test",
			},
		}
		getPermissionsRes = permissions.Permissions{
			UpdateBudget: true,
		}
	)

	tests := []struct {
		name          string
		accountFields model.UpdateReq
		wantError     error
		mockActions   func(interfaces)
	}{
		{
			"1. Редактирование счета (остаток и родительский счет) прошло успешно",
			updateReq,
			nil,
			func(inf interfaces) {
				// Успешно получили счет
				inf.accountRepository.On("Get", ctx, getAccountsReq).Return(getAccountRes, nil)
				// Получили разрешения
				inf.permissionsService.On("GetPermissions", getAccountRes[0]).Return(getPermissionsRes)
				// Проверка разрешений прошла успешно
				inf.permissionsService.On("CheckPermissions", updateReq, getPermissionsRes).Return(nil)
				// Успешно обновили остаток счета
				inf.accountService.On("ChangeRemainder", ctx, getAccountRes[0], *updateReq.Remainder).Return(nil)
				// Успешно проверили доступ к привязке к родительскому счету
				inf.accountService.On("ValidateUpdateParentAccountID", ctx, getAccountRes[0], *updateReq.ParentAccountID, updateReq.UserID).Return(nil)
				// Успешно обновили счет
				inf.accountRepository.On("Update", ctx, updateReq).Return(nil)
			},
		},
		{
			"2. Редактирование счета прошло успешно",
			model.UpdateReq{
				ID:     1,
				UserID: 1,
			},
			nil,
			func(inf interfaces) {
				updateReq := model.UpdateReq{
					ID:     1,
					UserID: 1,
				}
				// Успешно получили счет
				inf.accountRepository.On("Get", ctx, getAccountsReq).Return(getAccountRes, nil)
				// Получили разрешения
				inf.permissionsService.On("GetPermissions", getAccountRes[0]).Return(getPermissionsRes)
				// Проверка разрешений прошла успешно
				inf.permissionsService.On("CheckPermissions", updateReq, getPermissionsRes).Return(nil)
				// Успешно обновили счет
				inf.accountRepository.On("Update", ctx, updateReq).Return(nil)
			},
		},
		{
			"3. Редактирование счета прошло c ошибкой (остаток и родительский счет не изменялись)",
			model.UpdateReq{
				ID:     1,
				UserID: 1,
			},
			errors.InternalServer.New("Ошибка при обновлении счета"),
			func(inf interfaces) {
				updateReq := model.UpdateReq{
					ID:     1,
					UserID: 1,
				}
				// Успешно получили счет
				inf.accountRepository.On("Get", ctx, getAccountsReq).Return(getAccountRes, nil)
				// Получили разрешения
				inf.permissionsService.On("GetPermissions", getAccountRes[0]).Return(getPermissionsRes)
				// Проверка разрешений прошла успешно
				inf.permissionsService.On("CheckPermissions", updateReq, getPermissionsRes).Return(nil)
				// Получили ошибку при обновлении счета
				inf.accountRepository.On("Update", ctx, updateReq).Return(errors.InternalServer.New("Ошибка при обновлении счета"))
			},
		},
		{
			"4. Изменение баланса счета прошло c ошибкой",
			model.UpdateReq{
				ID:        1,
				UserID:    1,
				Remainder: pointer.Pointer(1.1),
			},
			errors.InternalServer.New("Ошибка при валидации счета"),
			func(inf interfaces) {
				updateReq := model.UpdateReq{
					ID:        1,
					UserID:    1,
					Remainder: pointer.Pointer(1.1),
				}
				// Успешно получили счет
				inf.accountRepository.On("Get", ctx, getAccountsReq).Return(getAccountRes, nil)
				// Получили разрешения
				inf.permissionsService.On("GetPermissions", getAccountRes[0]).Return(getPermissionsRes)
				// Проверка разрешений прошла успешно
				inf.permissionsService.On("CheckPermissions", updateReq, getPermissionsRes).Return(nil)
				// Обновление остатка прошло с ошибкой
				inf.accountService.On("ChangeRemainder", ctx, getAccountRes[0], *updateReq.Remainder).Return(errors.InternalServer.New("Ошибка при валидации счета"))
			},
		},
		{
			"5. Валидация родительского счета прошла c ошибкой",
			model.UpdateReq{
				ID:              1,
				UserID:          1,
				ParentAccountID: pointer.Pointer(uint32(1)),
			},
			errors.InternalServer.New("Ошибка при изменении баланса счета"),
			func(inf interfaces) {
				updateReq := model.UpdateReq{
					ID:              1,
					UserID:          1,
					ParentAccountID: pointer.Pointer(uint32(1)),
				}
				// Успешно получили счет
				inf.accountRepository.On("Get", ctx, getAccountsReq).Return(getAccountRes, nil)
				// Получили разрешения
				inf.permissionsService.On("GetPermissions", getAccountRes[0]).Return(getPermissionsRes)
				// Проверка разрешений прошла успешно
				inf.permissionsService.On("CheckPermissions", updateReq, getPermissionsRes).Return(nil)
				// Валидация родительского счета вернула ошибку
				inf.accountService.On("ValidateUpdateParentAccountID", ctx, getAccountRes[0], *updateReq.ParentAccountID, updateReq.UserID).Return(errors.InternalServer.New("Ошибка при валидации счета"))
			},
		},
		{
			"6. Действие пользователя запрещено",
			updateReq,
			errors.InternalServer.New("Ошибка при проверке разрешений"),
			func(inf interfaces) {
				// Успешно получили счет
				inf.accountRepository.On("Get", ctx, getAccountsReq).Return(getAccountRes, nil)
				// Получили разрешения
				inf.permissionsService.On("GetPermissions", getAccountRes[0]).Return(getPermissionsRes)
				// Проверка разрешений прошла с ошибкой
				inf.permissionsService.On("CheckPermissions", updateReq, getPermissionsRes).Return(errors.InternalServer.New("Ошибка при проверке разрешений"))
			},
		},
		{
			"7. Не найден счет для редактирования",
			updateReq,
			errors.NotFound.New("Счет не найден"),
			func(inf interfaces) {
				getAccountRes := []model.Account{}
				// Успешно получили счет
				inf.accountRepository.On("Get", ctx, getAccountsReq).Return(getAccountRes, nil)
			},
		},
		{
			"8. При получении счета произошла ошибка",
			updateReq,
			errors.InternalServer.New("Ошибка получения счета"),
			func(inf interfaces) {
				// Получили счет с ошибкой
				inf.accountRepository.On("Get", ctx, getAccountsReq).Return(nil, errors.InternalServer.New("Ошибка получения счета"))
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
			accountService := mocks.NewAccountService(t)

			s := &Service{
				accountRepository:  accountRepository,
				general:            generalRepository,
				transaction:        transactionRepository,
				user:               userRepository,
				permissionsService: permissionService,
				accountService:     accountService,
				logger:             logging.GetLogger(),
			}
			interfaces := interfaces{
				accountRepository:     accountRepository,
				transactionRepository: transactionRepository,
				generalRepository:     generalRepository,
				userRepository:        userRepository,
				permissionsService:    permissionService,
				accountService:        accountService,
			}

			tt.mockActions(interfaces)

			err := s.update(context.Background(), tt.accountFields)
			testingFunc.CheckError(t, err, tt.wantError)
		})
	}
}
