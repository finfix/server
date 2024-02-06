package converter

import (
	"core/app/proto/pbAccount"
	"jsonapi/app/internal/services/account/model"
)

type GetReq struct {
	model.GetReq
}

func (p GetReq) ConvertToProto() *pbAccount.GetReq {
	var pb pbAccount.GetReq
	pb.UserID = p.UserID
	pb.Accounting = p.Accounting
	pb.AccountGroupIDs = p.AccountGroupIDs
	pb.Type = p.Type.ConvertToOptionalProto()
	pb.DateFrom = p.DateFrom.ConvertToOptionalProto()
	pb.DateTo = p.DateTo.ConvertToOptionalProto()
	pb.Visible = p.Visible
	return &pb
}

type UpdateReq struct {
	model.UpdateReq
}

func (p UpdateReq) ConvertToProto() *pbAccount.UpdateReq {
	var pb pbAccount.UpdateReq
	pb.UserID = p.UserID
	pb.ID = p.ID
	pb.Remainder = p.Remainder
	pb.Name = p.Name
	pb.IconID = p.IconID
	pb.Visible = p.Visible
	pb.Accounting = p.Accounting
	pb.Budget = UpdateBudgetReq{p.Budget}.ConvertToProto()
	return &pb
}

type UpdateBudgetReq struct {
	model.UpdateBudgetReq
}

func (p UpdateBudgetReq) ConvertToProto() *pbAccount.UpdateBudgetReq {
	var pb pbAccount.UpdateBudgetReq
	pb.Amount = p.Amount
	pb.FixedSum = p.FixedSum
	pb.DaysOffset = p.DaysOffset
	pb.GradualFilling = p.GradualFilling
	return &pb
}

type CreateReq struct {
	model.CreateReq
}

func (p CreateReq) ConvertToProto() *pbAccount.CreateReq {
	var pb pbAccount.CreateReq
	pb.Budget = p.Budget
	pb.Remainder = p.Remainder
	pb.Name = p.Name
	pb.IconID = p.IconID
	pb.Type = p.Type.ConvertToProto()
	pb.Currency = p.Currency
	pb.AccountGroupID = p.AccountGroupID
	pb.Accounting = *p.Accounting
	pb.GradualBudgetFilling = *p.GradualBudgetFilling
	pb.UserID = p.UserID
	return &pb
}

type DeleteReq struct {
	model.DeleteReq
}

func (p DeleteReq) ConvertToProto() *pbAccount.DeleteReq {
	var pb pbAccount.DeleteReq
	pb.UserID = p.UserID
	pb.ID = p.ID
	return &pb
}

type SwitchReq struct {
	model.SwitchReq
}

func (p SwitchReq) ConvertToProto() *pbAccount.SwitchReq {
	var pb pbAccount.SwitchReq
	pb.UserID = p.UserID
	pb.ID1 = p.ID1
	pb.ID2 = p.ID2
	return &pb
}

type GetAccountGroupsReq struct {
	model.GetAccountGroupsReq
}

func (p GetAccountGroupsReq) ConvertToProto() *pbAccount.GetAccountGroupsReq {
	var pb pbAccount.GetAccountGroupsReq
	pb.UserID = p.UserID
	pb.AccountGroupIDs = p.AccountGroupIDs
	return &pb
}
