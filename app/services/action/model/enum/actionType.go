package enum

import (
	"server/app/pkg/errors"
)

type ActionType string

const (
	CreateTransaction = ActionType("transaction_create")
	UpdateTransaction = ActionType("transaction_update")
	DeleteTransaction = ActionType("transaction_delete")
	CreateAccount     = ActionType("account_create")
	UpdateAccount     = ActionType("account_update")
	DeleteAccount     = ActionType("account_delete")
	CreateUser        = ActionType("user_create")
	UpdateUser        = ActionType("user_update")
)

func (a *ActionType) Validate() error {
	switch *a {
	case CreateTransaction, UpdateTransaction, DeleteTransaction, CreateAccount, UpdateAccount, DeleteAccount, CreateUser, UpdateUser:
	default:
		return errors.BadRequest.New("Unknown action type",
			errors.StackTraceOption(errors.PreviousCaller),
			errors.ParamsOption("type", *a),
			errors.HumanTextOption("Неизвестный тип действия"),
		)
	}
	return nil
}
