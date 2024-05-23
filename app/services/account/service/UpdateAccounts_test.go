package service

import (
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"server/app/pkg/datetime"
	"server/app/pkg/log"
	"server/app/pkg/pointer"
	"server/app/pkg/testingFunc"
	"server/app/services/account/model"
	accountRepoModel "server/app/services/account/repository/model"
	mocks "server/mocks/server/app/services/account/service"
)

func TestService_update(t *testing.T) {

	log.Off()

	type interfaces struct {
		accountRepository     *mocks.MockAccountRepository
		transactionRepository *mocks.MockTransactionRepository
		generalRepository     *mocks.MockGeneralRepository
		userRepository        *mocks.MockUserRepository
		permissionsService    *mocks.MockAccountPermissionsService
		accountService        *mocks.MockAccountService
	}

	type args struct {
		ctx        context.Context
		account    model.Account
		updateReqs map[uint32]accountRepoModel.UpdateAccountReq
		userID     uint32
	}

	var (
		defaultArgs = args{
			ctx: context.Background(),
			account: model.Account{
				ID:                 1,
				Remainder:          decimal.Zero,
				Name:               "",
				IconID:             0,
				Type:               "",
				Currency:           "",
				Visible:            false,
				AccountGroupID:     0,
				AccountingInHeader: false,
				ParentAccountID:    nil,
				SerialNumber:       0,
				IsParent:           false,
				CreatedByUserID:    0,
				DatetimeCreate:     datetime.Time{Time: time.Now()},
				AccountingInCharts: false,
				AccountBudget: model.AccountBudget{
					Amount:         decimal.Zero,
					FixedSum:       decimal.Zero,
					DaysOffset:     0,
					GradualFilling: false,
				},
			},
			updateReqs: map[uint32]accountRepoModel.UpdateAccountReq{
				1: {
					Remainder:          pointer.Pointer(decimal.NewFromFloat(1.1)),
					Name:               nil,
					IconID:             nil,
					Visible:            nil,
					AccountingInHeader: nil,
					AccountingInCharts: nil,
					Currency:           nil,
					ParentAccountID:    nil,
					Budget: accountRepoModel.UpdateAccountBudgetReq{
						Amount:         nil,
						FixedSum:       nil,
						DaysOffset:     nil,
						GradualFilling: nil,
					},
					SerialNumber: nil,
				},
			},
			userID: 1,
		}
		updateRes = model.UpdateAccountRes{
			BalancingTransactionID:       pointer.Pointer(uint32(2)),
			BalancingAccountID:           pointer.Pointer(uint32(3)),
			BalancingAccountSerialNumber: pointer.Pointer(uint32(4)),
		}
	)

	tests := []struct {
		name        string
		args        args
		wantRes     model.UpdateAccountRes
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
				inf.accountService.On("ChangeAccountRemainder", defaultArgs.ctx, defaultArgs.account, *defaultArgs.updateReqs[1].Remainder, defaultArgs.userID).Return(updateRes, nil)
				// Успешно обновили счет
				inf.accountRepository.On("UpdateAccount", defaultArgs.ctx, defaultArgs.updateReqs).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountRepository := mocks.NewMockAccountRepository(t)
			generalRepository := mocks.NewMockGeneralRepository(t)
			transactionRepository := mocks.NewMockTransactionRepository(t)
			userRepository := mocks.NewMockUserRepository(t)
			accountPermissionService := mocks.NewMockAccountPermissionsService(t)
			accountService := mocks.NewMockAccountService(t)

			s := &Service{
				accountRepository:         accountRepository,
				general:                   generalRepository,
				transaction:               transactionRepository,
				user:                      userRepository,
				accountPermissionsService: accountPermissionService,
				accountService:            accountService,
			}
			interfaces := interfaces{
				accountRepository:     accountRepository,
				transactionRepository: transactionRepository,
				generalRepository:     generalRepository,
				userRepository:        userRepository,
				permissionsService:    accountPermissionService,
				accountService:        accountService,
			}

			tt.mockActions(interfaces)

			res, err := s.updateAccounts(tt.args.ctx, tt.args.account, tt.args.updateReqs, tt.args.userID)
			testingFunc.CheckError(t, err, tt.wantError)
			testingFunc.CheckStruct(t, res, tt.wantRes, nil)
		})
	}
}
