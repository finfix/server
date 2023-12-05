package model

import (
	"core/app/enum/transactionType"
	"core/app/proto/pbTransaction"
	"pkg/datetime/date"
	"pkg/datetime/time"
)

func (p *GetRes) ConvertToProto() *pbTransaction.GetRes {
	if p == nil {
		return nil
	}
	pb := &pbTransaction.GetRes{}
	pb.Transactions = make([]*pbTransaction.Transaction, len(p.Transactions))
	for i, transaction := range p.Transactions {
		pb.Transactions[i] = transaction.ConvertToProto()
	}
	return pb
}

func (p *Transaction) ConvertToProto() *pbTransaction.Transaction {
	if p == nil {
		return nil
	}
	pb := &pbTransaction.Transaction{}
	pb.ID = p.ID
	pb.Type = p.Type.ConvertToProto()
	pb.AmountFrom = p.AmountFrom
	pb.AmountTo = p.AmountTo
	pb.Note = p.Note
	pb.AccountFromID = p.AccountFromID
	pb.AccountToID = p.AccountToID
	pb.IsExecuted = p.IsExecuted
	pb.Accounting = p.Accounting
	pb.DateTransaction = p.DateTransaction.ConvertToProto()
	if p.TimeCreate != nil {
		pb.TimeCreate = time.Time{Time: *p.TimeCreate}.ConvertToProto()
	}
	return pb
}

func (p *Tag) ConvertToProto() *pbTransaction.Tag {
	if p == nil {
		return nil
	}
	pb := &pbTransaction.Tag{}
	pb.TransactionID = p.TransactionID
	pb.TagID = p.TagID
	return pb
}

func (p *CreateRes) ConvertToProto() *pbTransaction.CreateRes {
	if p == nil {
		return nil
	}
	pb := &pbTransaction.CreateRes{}
	pb.ID = p.ID
	return pb
}

type PbTag struct {
	*pbTransaction.Tag
}

func (pb PbTag) ConvertToStruct() *Tag {
	if pb.Tag == nil {
		return nil
	}
	var p Tag
	p.TransactionID = pb.TransactionID
	p.TagID = pb.TagID
	return &p
}

type PbGetReq struct {
	*pbTransaction.GetReq
}

func (pb PbGetReq) ConvertToStruct() *GetReq {
	if pb.GetReq == nil {
		return nil
	}
	var p GetReq
	p.UserID = pb.UserID
	p.AccountID = pb.AccountID
	p.Type = transactionType.PbTransactionType{pb.Type}.ConvertToOptionalEnum()
	p.DateFrom = date.PbDate{Timestamp: pb.DateFrom}.ConvertToOptionalDate()
	p.DateTo = date.PbDate{Timestamp: pb.DateTo}.ConvertToOptionalDate()
	p.Offset = pb.Offset
	p.Limit = pb.Limit
	return &p
}

type PbUpdateReq struct {
	*pbTransaction.UpdateReq
}

func (pb PbUpdateReq) ConvertToStruct() *UpdateReq {
	if pb.UpdateReq == nil {
		return nil
	}
	var p UpdateReq
	p.ID = pb.ID
	p.UserID = pb.UserID
	p.AmountFrom = pb.AmountFrom
	p.AmountTo = pb.AmountTo
	p.Note = pb.Note
	p.AccountFromID = pb.AccountFromID
	p.AccountToID = pb.AccountToID
	p.IsExecuted = pb.IsExecuted
	p.DateTransaction = date.PbDate{Timestamp: pb.DateTransaction}.ConvertToOptionalDate()
	return &p
}

type PbCreateReq struct {
	*pbTransaction.CreateReq
}

func (pb PbCreateReq) ConvertToStruct() *CreateReq {
	if pb.CreateReq == nil {
		return nil
	}
	var p CreateReq
	p.Type = transactionType.PbTransactionType{&pb.Type}.ConvertToEnum()
	p.AmountFrom = pb.AmountFrom
	p.AmountTo = pb.AmountTo
	p.Note = pb.Note
	p.AccountFromID = pb.AccountFromID
	p.AccountToID = pb.AccountToID
	p.DateTransaction = date.PbDate{pb.DateTransaction}.ConvertToDate()
	p.IsExecuted = pb.IsExecuted
	p.UserID = pb.UserID
	return &p
}

type PbDeleteReq struct {
	*pbTransaction.DeleteReq
}

func (pb PbDeleteReq) ConvertToStruct() *DeleteReq {
	if pb.DeleteReq == nil {
		return nil
	}
	var p DeleteReq
	p.UserID = pb.UserID
	p.ID = pb.ID
	return &p
}
