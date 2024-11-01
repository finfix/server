package transactionType

import (
	"pkg/errors"
)

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
		return errors.BadRequest.New("Unknown transaction type",
			errors.SkipThisCallOption(),
			errors.ParamsOption("type", *t),
			errors.HumanTextOption("Неизвестный тип транзакции"),
		)
	}
	return nil
}
