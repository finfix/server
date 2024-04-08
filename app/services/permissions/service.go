package permissions

import (
	"context"
	"time"

	model2 "server/app/services/account/model"
	"server/app/services/account/model/accountType"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/sql"
)

type Permissions struct {
	UpdateBudget    bool
	UpdateRemainder bool
	UpdateCurrency  bool

	DeleteAccount bool

	CreateTransaction bool

	LinkToParentAccount bool
}

type Service struct {
	db                    *sql.DB
	logger                *logging.Logger
	typeToPermissions     map[accountType.Type]Permissions
	isParentToPermissions map[bool]Permissions
}

var generalPermissions = Permissions{
	UpdateBudget:        true,
	UpdateRemainder:     true,
	UpdateCurrency:      true,
	DeleteAccount:       true,
	CreateTransaction:   true,
	LinkToParentAccount: true,
}

func (s *Service) GetPermissions(account model2.Account) Permissions {
	return joinPermissions(generalPermissions, s.typeToPermissions[account.Type], s.isParentToPermissions[account.IsParent])
}

func (s *Service) CheckPermissions(req model2.UpdateReq, permissions Permissions) error {

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
		joinedPermissions.UpdateCurrency = joinedPermissions.UpdateCurrency && permission.UpdateCurrency

		joinedPermissions.DeleteAccount = joinedPermissions.DeleteAccount && permission.DeleteAccount

		joinedPermissions.CreateTransaction = joinedPermissions.CreateTransaction && permission.CreateTransaction

		joinedPermissions.LinkToParentAccount = joinedPermissions.LinkToParentAccount && permission.LinkToParentAccount
	}
	return joinedPermissions
}

func (s *Service) refreshAccountPermissions(doOnce bool) error {
	for {
		var err error
		s.typeToPermissions, s.isParentToPermissions, err = s.getAccountPermissions(context.Background())
		if doOnce {
			return err
		}
		if err != nil {
			s.logger.Error(err)
		}

		time.Sleep(time.Minute)
	}
}

func (s *Service) getAccountPermissions(ctx context.Context) (
	_ map[accountType.Type]Permissions,
	_ map[bool]Permissions,
	err error,
) {

	rows, err := s.db.Query(ctx, `
		SELECT * 
		FROM account_permissions`)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	typeToPermissions := make(map[accountType.Type]Permissions)
	isParentToPermissions := make(map[bool]Permissions)

	for rows.Next() {
		var _accountType, actionType string
		var access bool
		if err := rows.Scan(&_accountType, &actionType, &access); err != nil {
			return nil, nil, err
		}
		permission := Permissions{
			UpdateBudget:        actionType == "updateBudget" && access,
			UpdateRemainder:     actionType == "updateRemainder" && access,
			UpdateCurrency:      actionType == "updateCurrency" && access,
			DeleteAccount:       actionType == "deleteAccount" && access,
			CreateTransaction:   actionType == "createTransaction" && access,
			LinkToParentAccount: actionType == "linkToParentAccount" && access,
		}
		switch _accountType {
		case "regular", "debt", "earnings", "expense":
			typeToPermissions[accountType.Type(_accountType)] = permission
		case "parent", "general":
			isParentToPermissions[_accountType == "parent"] = permission
		}
	}

	return typeToPermissions, isParentToPermissions, nil
}

func New(
	db *sql.DB,
	logger *logging.Logger,
) (*Service, error) {

	service := &Service{
		db:     db,
		logger: logger,
	}

	logger.Info("Получаем пермишены на действия со счетами")
	err := service.refreshAccountPermissions(true)
	if err != nil {
		return nil, err
	}
	go service.refreshAccountPermissions(false)

	return service, nil
}
