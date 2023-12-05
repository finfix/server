package accountType

import (
	"core/app/proto/pbEnum"
	"pkg/errors"
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

type PbAccountType struct {
	*pbEnum.AccountType
}

func (pb PbAccountType) ConvertToEnum() Type {
	if pb.AccountType == nil {
		return ""
	}
	switch *pb.AccountType {
	case pbEnum.AccountType_Regular:
		return Regular
	case pbEnum.AccountType_Expense:
		return Expense
	case pbEnum.AccountType_Debt:
		return Debt
	case pbEnum.AccountType_Earnings:
		return Earnings
	}
	return ""
}

func (pb PbAccountType) ConvertToOptionalEnum() *Type {
	if pb.AccountType == nil {
		return nil
	}
	t := pb.ConvertToEnum()
	return &t
}

func (t Type) ConvertToProto() pbEnum.AccountType {
	switch t {
	case Earnings:
		return pbEnum.AccountType_Earnings
	case Expense:
		return pbEnum.AccountType_Expense
	case Debt:
		return pbEnum.AccountType_Debt
	case Regular:
		return pbEnum.AccountType_Regular
	}
	return pbEnum.AccountType_Regular
}

func (t *Type) ConvertToOptionalProto() *pbEnum.AccountType {
	if t == nil {
		return nil
	}
	pb := t.ConvertToProto()
	return &pb
}
