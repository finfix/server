package accountType

import (
	"server/pkg/errors"
)

const stackDepth = 2

type Type string

// enums:"regular,expense,debt,income,balancing"
const (
	Regular   = Type("regular")
	Expense   = Type("expense")
	Debt      = Type("debt")
	Earnings  = Type("earnings")
	Balancing = Type("balancing")
)

func (t *Type) Validate() error {
	if t == nil {
		return nil
	}
	switch *t {
	case Earnings, Expense, Debt, Regular, Balancing:
	default:
		err := errors.BadRequest.NewPathCtx("Unknown account type", stackDepth, "type: %v", *t)
		return errors.AddHumanText(err, "Неизвестный тип счета")
	}
	return nil
}
