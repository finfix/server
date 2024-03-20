package transactionType

import (
	"server/pkg/errors"
)

const stackDepth = 2

type Type string

// enum:"consumption,income,transfer"
const (
	Transfer    = Type("transfer")
	Consumption = Type("consumption")
	Balancing   = Type("balancing")
	Income      = Type("income")
)

func (t *Type) Validate() error {
	if t == nil {
		return nil
	}
	switch *t {
	case Transfer, Consumption, Balancing, Income:
	default:
		err := errors.BadRequest.NewPathCtx("Unknown transaction type", stackDepth, "type: %v", *t)
		return errors.AddHumanText(err, "Неизвестный тип транзакции")
	}
	return nil
}
