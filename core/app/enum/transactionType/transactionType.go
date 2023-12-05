package transactionType

import (
	"core/app/proto/pbEnum"
	"pkg/errors"
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

type PbTransactionType struct {
	*pbEnum.TransactionType
}

func (pb PbTransactionType) ConvertToEnum() Type {
	if pb.TransactionType == nil {
		return ""
	}
	switch *pb.TransactionType {
	case pbEnum.TransactionType_Transfer:
		return Transfer
	case pbEnum.TransactionType_Consumption:
		return Consumption
	case pbEnum.TransactionType_Balancing:
		return Balancing
	case pbEnum.TransactionType_Income:
		return Income
	}
	return ""
}

func (pb PbTransactionType) ConvertToOptionalEnum() *Type {
	if pb.TransactionType == nil {
		return nil
	}
	transactionType := pb.ConvertToEnum()
	return &transactionType
}

func (t Type) ConvertToProto() pbEnum.TransactionType {
	switch t {
	case Transfer:
		return pbEnum.TransactionType_Transfer
	case Consumption:
		return pbEnum.TransactionType_Consumption
	case Balancing:
		return pbEnum.TransactionType_Balancing
	case Income:
		return pbEnum.TransactionType_Income
	}
	return pbEnum.TransactionType_Consumption
}

func (t *Type) ConvertToOptionalProto() *pbEnum.TransactionType {
	if t == nil {
		return nil
	}
	transactionType := t.ConvertToProto()
	return &transactionType
}
