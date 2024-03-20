package accountType

import (
	"server/pkg/errors"
)

const stackDepth = 2

type Type string

// enum:"regular,expense,credit,debt,income,investments"
const (
	Regular  = Type("regular")
	Expense  = Type("expense")
	Debt     = Type("debt")
	Earnings = Type("earnings")
)

func (t *Type) Validate() error {
	if t == nil {
		return nil
	}
	switch *t {
	case Earnings, Expense, Debt, Regular:
	default:
		err := errors.BadRequest.NewPathCtx("Unknown account type", stackDepth, "type: %v", *t)
		return errors.AddHumanText(err, "Неизвестный тип счета")
	}
	return nil
}
