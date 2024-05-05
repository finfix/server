package accountPermissions

import (
	"context"
	"sync"
	"time"

	"server/app/pkg/errors"
	"server/app/pkg/log"
	"server/app/pkg/sql"
	"server/app/services/account/model"
	"server/app/services/account/model/accountType"
)

type AccountPermissions struct {
	UpdateBudget          bool
	UpdateRemainder       bool
	UpdateCurrency        bool
	UpdateParentAccountID bool

	CreateTransaction bool
}

type Service struct {
	db sql.SQL

	permissions permissions
}

type permissions struct {
	typeToPermissions     map[accountType.Type]AccountPermissions
	isParentToPermissions map[bool]AccountPermissions
	mu                    sync.RWMutex
}

func (p *permissions) get() (
	map[accountType.Type]AccountPermissions,
	map[bool]AccountPermissions,
) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.typeToPermissions, p.isParentToPermissions
}

func (p *permissions) set(
	typeToPermissions map[accountType.Type]AccountPermissions,
	isParentToPermissions map[bool]AccountPermissions,
) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.typeToPermissions = typeToPermissions
	p.isParentToPermissions = isParentToPermissions
}

var generalPermissions = AccountPermissions{
	UpdateBudget:          true,
	UpdateRemainder:       true,
	UpdateCurrency:        true,
	UpdateParentAccountID: true,

	CreateTransaction: true,
}

func (s *Service) GetAccountPermissions(account model.Account) AccountPermissions {
	typeToPermissions, isParentToPermissions := s.permissions.get()
	return joinAccountPermissions(
		generalPermissions,
		typeToPermissions[account.Type],
		isParentToPermissions[account.IsParent],
	)
}

func (s *Service) CheckAccountPermissions(req model.UpdateAccountReq, permissions AccountPermissions) error {

	if (req.Budget.DaysOffset != nil || req.Budget.Amount != nil || req.Budget.FixedSum != nil || req.Budget.GradualFilling != nil) && !permissions.UpdateBudget {
		return errors.BadRequest.New("Нельзя менять бюджет")
	}

	if req.Currency != nil && !permissions.UpdateCurrency {
		return errors.BadRequest.New("Нельзя менять валюту")
	}

	if req.Remainder != nil && !permissions.UpdateRemainder {
		return errors.BadRequest.New("Нельзя менять остаток")
	}

	return nil
}

func joinAccountPermissions(permissions ...AccountPermissions) (joinedPermissions AccountPermissions) {
	joinedPermissions = generalPermissions
	for _, permission := range permissions {
		joinedPermissions.UpdateBudget = joinedPermissions.UpdateBudget && permission.UpdateBudget
		joinedPermissions.UpdateRemainder = joinedPermissions.UpdateRemainder && permission.UpdateRemainder
		joinedPermissions.UpdateCurrency = joinedPermissions.UpdateCurrency && permission.UpdateCurrency
		joinedPermissions.UpdateParentAccountID = joinedPermissions.UpdateParentAccountID && permission.UpdateParentAccountID

		joinedPermissions.CreateTransaction = joinedPermissions.CreateTransaction && permission.CreateTransaction
	}
	return joinedPermissions
}

func (s *Service) refreshAccountPermissions(doOnce bool) error {
	for {
		var err error
		_typeToPermissions, _isParentToPermissions, err := s.getAccountPermissions(context.Background())
		s.permissions.set(_typeToPermissions, _isParentToPermissions)
		if doOnce {
			return err
		}
		if err != nil {
			log.Error(context.Background(), err)
		}

		time.Sleep(time.Minute)
	}
}

func (s *Service) getAccountPermissions(ctx context.Context) (
	_ map[accountType.Type]AccountPermissions,
	_ map[bool]AccountPermissions,
	err error,
) {

	rows, err := s.db.Query(ctx, `
		SELECT * 
		FROM permissions.account_permissions`)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	typeToPermissions := make(map[accountType.Type]AccountPermissions)
	isParentToPermissions := make(map[bool]AccountPermissions)

	for rows.Next() {
		var _accountType, actionType string
		var access bool
		if err := rows.Scan(&_accountType, &actionType, &access); err != nil {
			return nil, nil, err
		}

		var permission AccountPermissions
		switch _accountType {
		case "regular", "debt", "earnings", "expense", "balancing":
			permission = typeToPermissions[accountType.Type(_accountType)]
		case "parent", "general": //nolint:goconst
			permission = isParentToPermissions[_accountType == "parent"] //nolint:goconst
		}

		switch actionType {
		case "update_budget":
			permission.UpdateBudget = access
		case "update_remainder":
			permission.UpdateRemainder = access
		case "update_currency":
			permission.UpdateCurrency = access
		case "update_parent_account_id":
			permission.UpdateParentAccountID = access
		case "create_transaction":
			permission.CreateTransaction = access
		}

		switch _accountType {
		case "regular", "debt", "earnings", "expense", "balancing":
			typeToPermissions[accountType.Type(_accountType)] = permission
		case "parent", "general":
			isParentToPermissions[_accountType == "parent"] = permission
		}
	}

	return typeToPermissions, isParentToPermissions, nil
}

func New(
	db sql.SQL,

) (*Service, error) {

	service := &Service{
		db: db,
		permissions: permissions{
			typeToPermissions:     nil,
			isParentToPermissions: nil,
			mu:                    sync.RWMutex{},
		},
	}

	log.Info(context.Background(), "Получаем пермишены на действия со счетами")
	err := service.refreshAccountPermissions(true)
	if err != nil {
		return nil, err
	}
	go func() {
		_ = service.refreshAccountPermissions(false)
	}()

	return service, nil
}
