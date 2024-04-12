package service

import (
	"context"
	"testing"

	"server/app/pkg/errors"
	"server/app/pkg/logging"
	"server/app/pkg/pointer"
	"server/app/pkg/testingFunc"
	"server/app/services/account/model"
	accountRepoModel "server/app/services/account/repository/model"
	mocks "server/mocks/server/app/services/account/service"
)

func TestService_update(t *testing.T) {

	logging.Off()

	type interfaces struct {
		accountRepository     *mocks.MockAccountRepository
		transactionRepository *mocks.MockTransactionRepository
		generalRepository     *mocks.MockGeneralRepository
		userRepository        *mocks.MockUserRepository
		permissionsService    *mocks.MockPermissionsService
		accountService        *mocks.MockAccountService
	}

	type args struct {
		ctx        context.Context
		account    model.Account
		updateReqs map[uint32]accountRepoModel.UpdateReq
	}

	var (
		defaultArgs = args{
			ctx: context.Background(),
			account: model.Account{
				ID: 1,
			},
			updateReqs: map[uint32]accountRepoModel.UpdateReq{
				1: {
					Remainder: pointer.Pointer(1.1),
				},
			},
		}
		updateRes = model.UpdateRes{
			BalancingTransactionID:       pointer.Pointer(uint32(2)),
			BalancingAccountID:           pointer.Pointer(uint32(3)),
			BalancingAccountSerialNumber: pointer.Pointer(uint32(4)),
		}
	)

	tests := []struct {
		name        string
		args        args
		wantRes     model.UpdateRes
		wantError   error
		mockActions func(interfaces)
	}{
		{
			"1. Редактирование счета (с изменением остатка) прошло успешно",
			defaultArgs,
			updateRes,
			nil,
			func(inf interfaces) {
				// Успешно обновили остаток счета
				inf.accountService.On("ChangeRemainder", defaultArgs.ctx, defaultArgs.account, *defaultArgs.updateReqs[1].Remainder).Return(updateRes, nil)
				// Успешно обновили счет
				inf.accountRepository.On("Update", defaultArgs.ctx, defaultArgs.updateReqs).Return(nil)
			},
		},
		{
			"2. Редактирование счета (без изменения остатка) прошло успешно",
			args{
				ctx:     defaultArgs.ctx,
				account: defaultArgs.account,
				updateReqs: map[uint32]accountRepoModel.UpdateReq{
					1: {
						Remainder: nil,
					},
				},
			},
			model.UpdateRes{},
			nil,
			func(inf interfaces) {
				updateReqs := map[uint32]accountRepoModel.UpdateReq{
					1: {
						Remainder: nil,
					},
				}
				// Успешно обновили счет
				inf.accountRepository.On("Update", defaultArgs.ctx, updateReqs).Return(nil)
			},
		},
		{
			"3. Редактирование счета (без изменения остатка) прошло c ошибкой",
			args{
				ctx:     defaultArgs.ctx,
				account: defaultArgs.account,
				updateReqs: map[uint32]accountRepoModel.UpdateReq{
					1: {
						Remainder: nil,
					},
				},
			},
			model.UpdateRes{},
			errors.InternalServer.New("Ошибка при обновлении счета"),
			func(inf interfaces) {
				updateReqs := map[uint32]accountRepoModel.UpdateReq{
					1: {
						Remainder: nil,
					},
				}
				// Получили ошибку при обновлении счета
				inf.accountRepository.On("Update", defaultArgs.ctx, updateReqs).Return(errors.InternalServer.New("Ошибка при обновлении счета"))
			},
		},
		{
			"4. Изменение баланса счета прошло c ошибкой",
			defaultArgs,
			model.UpdateRes{},
			errors.InternalServer.New("Ошибка при валидации счета"),
			func(inf interfaces) {
				// Обновление остатка прошло с ошибкой
				inf.accountService.On("ChangeRemainder", defaultArgs.ctx, defaultArgs.account, *defaultArgs.updateReqs[defaultArgs.account.ID].Remainder).Return(model.UpdateRes{}, errors.InternalServer.New("Ошибка при валидации счета"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountRepository := mocks.NewMockAccountRepository(t)
			generalRepository := mocks.NewMockGeneralRepository(t)
			transactionRepository := mocks.NewMockTransactionRepository(t)
			userRepository := mocks.NewMockUserRepository(t)
			permissionService := mocks.NewMockPermissionsService(t)
			accountService := mocks.NewMockAccountService(t)

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

			err, res := s.update(tt.args.ctx, tt.args.account, tt.args.updateReqs)
			testingFunc.CheckError(t, err, tt.wantError)
			testingFunc.CheckStruct(t, res, tt.wantRes, nil)
		})
	}
}
