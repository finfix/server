package enum

import (
	"server/pkg/errors"
)

const stackDepth = 2

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
		return errors.BadRequest.New("Unknown action type", errors.Options{
			Params:    map[string]any{"type": *a},
			HumanText: "Неизвестный тип действия",
		})
	}
	return nil
}
