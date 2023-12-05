package converter

import (
	"core/app/internal/services/account/model"
	"core/app/proto/pbAccount"
)

type CreateRes struct {
	model.CreateRes
}

func (p CreateRes) ConvertToProto() *pbAccount.CreateRes {
	var pb pbAccount.CreateRes
	pb.ID = p.ID
	return &pb
}

type GetRes struct {
	model.GetRes
}

func (p GetRes) ConvertToProto() *pbAccount.GetRes {
	var pb pbAccount.GetRes
	pb.Accounts = make([]*pbAccount.Account, len(p.Accounts))
	for i, account := range p.Accounts {
		pb.Accounts[i] = Account{account}.ConvertToProto()
	}
	return &pb
}

type Account struct {
	model.Account
}

func (p Account) ConvertToProto() *pbAccount.Account {
	var pb pbAccount.Account
	pb.ID = p.ID
	pb.Budget = p.Budget
	pb.Remainder = p.Remainder
	pb.Name = p.Name
	pb.IconID = p.IconID
	pb.Type = p.Type.ConvertToProto()
	pb.Currency = p.Currency
	pb.Visible = p.Visible
	pb.AccountGroupID = p.AccountGroupID
	pb.Accounting = p.Accounting
	pb.ParentAccountID = p.ParentAccountID
	pb.GradualBudgetFilling = p.GradualBudgetFilling
	pb.SerialNumber = p.SerialNumber
	return &pb
}

type QuickStatisticRes struct {
	model.QuickStatisticRes
}

func (p QuickStatisticRes) ConvertToProto() *pbAccount.QuickStatisticRes {
	var pb pbAccount.QuickStatisticRes
	pb.QuickStatistic = make([]*pbAccount.QuickStatistic, len(p.QuickStatisticRes.QuickStatistic))
	for i, statistic := range p.QuickStatistic {
		pb.QuickStatistic[i] = QuickStatistic{statistic}.ConvertToProto()
	}
	return &pb
}

type QuickStatistic struct {
	model.QuickStatistic
}

func (p QuickStatistic) ConvertToProto() *pbAccount.QuickStatistic {
	var pb pbAccount.QuickStatistic
	pb.TotalRemainder = p.TotalRemainder
	pb.TotalExpense = p.TotalExpense
	pb.AccountGroupID = p.AccountGroupID
	pb.TotalBudget = p.TotalBudget
	pb.Currency = p.Currency
	return &pb
}

type GetAccountGroupsRes struct {
	model.GetAccountGroupsRes
}

func (p GetAccountGroupsRes) ConvertToProto() *pbAccount.GetAccountGroupsRes {
	var pb pbAccount.GetAccountGroupsRes
	pb.AccountGroups = make([]*pbAccount.AccountGroup, len(p.AccountGroups))
	for i, accountGroup := range p.AccountGroups {
		pb.AccountGroups[i] = AccountGroup{accountGroup}.ConvertToProto()
	}
	return &pb
}

type AccountGroup struct {
	model.AccountGroup
}

func (p AccountGroup) ConvertToProto() *pbAccount.AccountGroup {
	var pb pbAccount.AccountGroup
	pb.ID = p.ID
	pb.Name = p.Name
	pb.Currency = p.Currency
	pb.SerialNumber = p.SerialNumber
	pb.Visible = p.Visible
	return &pb
}
