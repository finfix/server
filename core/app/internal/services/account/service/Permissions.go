package service

import (
	"context"
	"core/app/enum/accountType"
	"core/app/internal/services/account/model"
	"pkg/errors"
)

type Permissions struct {
	UpdateBudget bool

	UpdateRemainder bool

	DeleteAccount bool

	CreateTransaction bool

	LinkToParentAccount bool
}

func (s *Service) GetPermissions(ctx context.Context, accountID uint32) (permissions Permissions, err error) {
	accounts, err := s.account.Get(ctx, model.GetReq{IDs: []uint32{accountID}})
	if err != nil {
		return permissions, err
	}
	if len(accounts) == 0 {
		return permissions, errors.NotFound.New("Аккаунт не найден")
	}
	account := accounts[0]

	return joinPermissions(generalPermissions, typeToPermissions[account.Type], isParentToPermissions[account.IsParent]), nil
}

func (s *Service) CheckPermissions(req model.UpdateReq, permissions Permissions) error {

	if (req.Budget.DaysOffset != nil || req.Budget.Amount != nil || req.Budget.FixedSum != nil || req.Budget.GradualFilling != nil) && !permissions.UpdateBudget {
		return errors.BadRequest.New("Нельзя менять бюджет")
	}

	if req.Remainder != nil && !permissions.UpdateRemainder {
		return errors.BadRequest.New("Нельзя менять остаток")
	}

	return nil
}

func joinPermissions(permissions ...Permissions) (joinedPermissions Permissions) {
	joinedPermissions = generalPermissions
	for _, permission := range permissions {
		joinedPermissions.UpdateBudget = joinedPermissions.UpdateBudget && permission.UpdateBudget
		joinedPermissions.UpdateRemainder = joinedPermissions.UpdateRemainder && permission.UpdateRemainder

		joinedPermissions.DeleteAccount = joinedPermissions.DeleteAccount && permission.DeleteAccount

		joinedPermissions.CreateTransaction = joinedPermissions.CreateTransaction && permission.CreateTransaction

		joinedPermissions.LinkToParentAccount = joinedPermissions.LinkToParentAccount && permission.LinkToParentAccount
	}
	return joinedPermissions
}

var (
	typeToPermissions = map[accountType.Type]Permissions{
		accountType.Regular:  regularPermissions,
		accountType.Debt:     debtPermissions,
		accountType.Earnings: earningsPermissions,
		accountType.Expense:  expensePermissions,
	}
	isParentToPermissions = map[bool]Permissions{
		true:  parentPermissions,
		false: generalPermissions,
	}
)

var (
	// Общие разрешения для всех типов счетов
	generalPermissions = Permissions{
		UpdateBudget:        true,
		UpdateRemainder:     true,
		DeleteAccount:       true,
		CreateTransaction:   true,
		LinkToParentAccount: true,
	}

	// Разрешения для обычных счетов - нельзя менять только бюджеты
	regularPermissions = Permissions{
		UpdateBudget:        false,
		UpdateRemainder:     true,
		DeleteAccount:       true,
		CreateTransaction:   true,
		LinkToParentAccount: true,
	}

	// Разрешения для счетов долгов - нельзя менять только бюджеты
	debtPermissions = Permissions{
		UpdateBudget:        false,
		UpdateRemainder:     true,
		DeleteAccount:       true,
		CreateTransaction:   true,
		LinkToParentAccount: true,
	}

	// Разрешения для счетов доходов - нельзя менять только остатки счетов
	earningsPermissions = Permissions{
		UpdateBudget:        true,
		UpdateRemainder:     false,
		DeleteAccount:       true,
		CreateTransaction:   true,
		LinkToParentAccount: true,
	}

	// Разрешения для счетов расходов - нельзя менять только остатки счетов
	expensePermissions = Permissions{
		UpdateBudget:        true,
		UpdateRemainder:     false,
		DeleteAccount:       true,
		CreateTransaction:   true,
		LinkToParentAccount: true,
	}

	// Разрешения для родительских счетов - нельзя менять остатки счетов, создавать транзакции и привязывать к родительским счетам
	parentPermissions = Permissions{
		UpdateBudget:        true,
		UpdateRemainder:     false,
		DeleteAccount:       true,
		CreateTransaction:   false,
		LinkToParentAccount: false,
	}
)
