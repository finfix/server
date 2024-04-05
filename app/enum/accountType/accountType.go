package accountType

import (
	"server/pkg/errors"
)

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
		return errors.BadRequest.New("Unknown account type", errors.Options{
			PathDepth: errors.SecondPathDepth,
			Params:    map[string]any{"type": *t},
			HumanText: "Неизвестный тип счета",
		})
	}
	return nil
}
