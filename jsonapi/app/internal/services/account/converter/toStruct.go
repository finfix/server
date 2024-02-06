package converter

import (
	"core/app/enum/accountType"
	"core/app/proto/pbAccount"
	"jsonapi/app/internal/services/account/model"
)

type PbGetRes struct {
	*pbAccount.GetRes
}

func (pb PbGetRes) ConvertToStruct() model.GetRes {
	var p model.GetRes
	p.Accounts = make([]model.Account, len(pb.Accounts))
	for i, account := range pb.Accounts {
		p.Accounts[i] = PbAccount{account}.ConvertToStruct()
	}
	return p
}

type PbAccount struct {
	*pbAccount.Account
}

func (pb PbAccount) ConvertToStruct() model.Account {
	var p model.Account
	p.ID = pb.ID
	p.Budget = pb.Budget
	p.Remainder = pb.Remainder
	p.Name = pb.Name
	p.IconID = pb.IconID
	p.Type = accountType.PbAccountType{&pb.Type}.ConvertToEnum()
	p.Currency = pb.Currency
	p.Visible = pb.Visible
	p.AccountGroupID = pb.AccountGroupID
	p.Accounting = pb.Accounting
	p.ParentAccountID = pb.ParentAccountID
	p.GradualBudgetFilling = pb.GradualBudgetFilling
	p.SerialNumber = pb.SerialNumber
	p.IsParent = pb.IsParent
	return p
}

type PbCreateRes struct {
	*pbAccount.CreateRes
}

func (pb PbCreateRes) ConvertToStruct() model.CreateRes {
	var p model.CreateRes
	p.ID = pb.ID
	return p
}

type PbQuickStatistic struct {
	*pbAccount.QuickStatistic
}

func (pb PbBudget) ConvertToStruct() model.Budget {
	var p model.Budget
	p.Amount = pb.Amount
	p.FixedSum = pb.FixedSum
	p.DaysOffset = pb.DaysOffset
	p.GradualFilling = pb.GradualFilling
	return p
}

type PbCreateRes struct {
	*pbAccount.CreateRes
}

func (pb PbCreateRes) ConvertToStruct() model.CreateRes {
	var p model.CreateRes
	p.ID = pb.ID
	return p
}

type PbGetAccountGroupsRes struct {
	*pbAccount.GetAccountGroupsRes
}

func (pb PbGetAccountGroupsRes) ConvertToStruct() model.GetAccountGroupsRes {
	var p model.GetAccountGroupsRes
	p.AccountGroups = make([]model.AccountGroup, len(pb.AccountGroups))
	for i, accountGroup := range pb.AccountGroups {
		p.AccountGroups[i] = PbAccountGroup{accountGroup}.ConvertToStruct()
	}
	return p
}

type PbAccountGroup struct {
	*pbAccount.AccountGroup
}

func (pb PbAccountGroup) ConvertToStruct() model.AccountGroup {
	var p model.AccountGroup
	p.ID = pb.ID
	p.Name = pb.Name
	p.Currency = pb.Currency
	p.SerialNumber = pb.SerialNumber
	p.Visible = pb.Visible
	return p
}
