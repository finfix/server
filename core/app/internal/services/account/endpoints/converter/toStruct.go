package converter

import (
	"core/app/enum/accountType"
	"core/app/internal/services/account/model"
	"core/app/proto/pbAccount"
	"pkg/datetime/date"
)

type PbGetReq struct {
	*pbAccount.GetReq
}

func (pb PbGetReq) ConvertToStruct() model.GetReq {
	var p model.GetReq
	p.UserID = pb.UserID
	p.Accounting = pb.Accounting
	p.Visible = pb.Visible
	p.AccountGroupIDs = pb.AccountGroupIDs
	p.Type = accountType.PbAccountType{AccountType: pb.Type}.ConvertToOptionalEnum()
	p.DateFrom = date.PbDate{Timestamp: pb.DateFrom}.ConvertToOptionalDate()
	p.DateTo = date.PbDate{Timestamp: pb.DateTo}.ConvertToOptionalDate()
	return p
}

type PbUpdateReq struct {
	*pbAccount.UpdateReq
}

func (pb PbUpdateReq) ConvertToStruct() model.UpdateReq {
	var p model.UpdateReq
	p.UserID = pb.UserID
	p.ID = pb.ID
	p.Budget = pb.Budget
	p.Remainder = pb.Remainder
	p.Name = pb.Name
	p.IconID = pb.IconID
	p.Visible = pb.Visible
	p.AccountGroupID = pb.AccountGroupID
	p.Accounting = pb.Accounting
	p.GradualBudgetFilling = pb.GradualBudgetFilling
	return p
}

type PbCreateReq struct {
	*pbAccount.CreateReq
}

func (pb PbCreateReq) ConvertToStruct() model.CreateReq {
	var p model.CreateReq
	p.Budget = pb.Budget
	p.Remainder = pb.Remainder
	p.Name = pb.Name
	p.IconID = pb.IconID
	p.Type = accountType.PbAccountType{&pb.Type}.ConvertToEnum()
	p.Currency = pb.Currency
	p.AccountGroupID = pb.AccountGroupID
	p.Accounting = pb.Accounting
	p.GradualBudgetFilling = pb.GradualBudgetFilling
	p.UserID = pb.UserID
	return p
}

type PbDeleteReq struct {
	*pbAccount.DeleteReq
}

func (pb PbDeleteReq) ConvertToStruct() model.DeleteReq {
	var p model.DeleteReq
	p.UserID = pb.UserID
	p.ID = pb.ID
	return p
}

type PbSwitchReq struct {
	*pbAccount.SwitchReq
}

func (pb PbSwitchReq) ConvertToStruct() model.SwitchReq {
	var p model.SwitchReq
	p.UserID = pb.UserID
	p.ID1 = pb.ID1
	p.ID2 = pb.ID2
	return p
}

type PbQuickStatisticReq struct {
	*pbAccount.QuickStatisticReq
}

func (pb PbQuickStatisticReq) ConvertToStruct() model.QuickStatisticReq {
	var p model.QuickStatisticReq
	p.UserID = pb.UserID
	return p
}

type PbGetAccountGroupsReq struct {
	*pbAccount.GetAccountGroupsReq
}

func (pb PbGetAccountGroupsReq) ConvertToStruct() model.GetAccountGroupsReq {
	var p model.GetAccountGroupsReq
	p.UserID = pb.UserID
	p.AccountGroupIDs = pb.AccountGroupIDs
	return p
}
