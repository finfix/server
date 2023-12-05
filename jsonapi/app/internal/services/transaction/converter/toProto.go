package converter

import (
	"core/app/proto/pbTransaction"
	"jsonapi/app/internal/services/transaction/model"
)

type GetReq struct {
	model.GetReq
}

func (p GetReq) ConvertToProto() *pbTransaction.GetReq {
	var pb pbTransaction.GetReq
	pb.UserID = p.UserID
	pb.Offset = p.Offset
	pb.Limit = p.Limit
	pb.AccountID = p.AccountID
	pb.Type = p.Type.ConvertToOptionalProto()
	pb.DateFrom = p.DateFrom.ConvertToOptionalProto()
	pb.DateTo = p.DateTo.ConvertToOptionalProto()
	return &pb
}

type UpdateReq struct {
	model.UpdateReq
}

func (p UpdateReq) ConvertToProto() *pbTransaction.UpdateReq {
	var pb pbTransaction.UpdateReq
	pb.ID = p.ID
	pb.UserID = p.UserID
	pb.AmountFrom = p.AmountFrom
	pb.AmountTo = p.AmountTo
	pb.Note = p.Note
	pb.AccountFromID = p.AccountFromID
	pb.AccountToID = p.AccountToID
	pb.IsExecuted = p.IsExecuted
	pb.DateTransaction = p.DateTransaction.ConvertToOptionalProto()
	return &pb
}

type CreateReq struct {
	model.CreateReq
}

func (p CreateReq) ConvertToProto() *pbTransaction.CreateReq {
	var pb pbTransaction.CreateReq
	pb.Type = p.Type.ConvertToProto()
	pb.AmountFrom = p.AmountFrom
	pb.AmountTo = p.AmountTo
	pb.Note = p.Note
	pb.AccountFromID = p.AccountFromID
	pb.AccountToID = p.AccountToID
	pb.DateTransaction = p.DateTransaction.ConvertToProto()
	pb.IsExecuted = p.IsExecuted
	pb.UserID = p.UserID
	return &pb
}

type DeleteReq struct {
	model.DeleteReq
}

func (p DeleteReq) ConvertToProto() *pbTransaction.DeleteReq {
	var pb pbTransaction.DeleteReq
	pb.UserID = p.UserID
	pb.ID = p.ID
	return &pb
}
