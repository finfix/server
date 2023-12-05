package converter

import (
	"core/app/enum/transactionType"
	"core/app/proto/pbTransaction"
	"jsonapi/app/internal/services/transaction/model"
	"pkg/datetime/date"
	"pkg/datetime/time"
)

type PbGetRes struct {
	*pbTransaction.GetRes
}

func (pb PbGetRes) ConvertToStruct() model.GetRes {
	var p model.GetRes
	p.Transactions = make([]model.Transaction, len(pb.Transactions))
	for i, transaction := range pb.Transactions {
		p.Transactions[i] = PbTransaction{transaction}.ConvertToStruct()
	}
	return p
}

type PbTransaction struct {
	*pbTransaction.Transaction
}

func (pb PbTransaction) ConvertToStruct() model.Transaction {
	var p model.Transaction
	p.ID = pb.ID
	p.Type = transactionType.PbTransactionType{&pb.Type}.ConvertToEnum()
	p.AmountFrom = pb.AmountFrom
	p.AmountTo = pb.AmountTo
	p.Note = pb.Note
	p.AccountFromID = pb.AccountFromID
	p.AccountToID = pb.AccountToID
	p.IsExecuted = pb.IsExecuted
	p.Accounting = pb.Accounting
	p.TimeCreate = time.PbTime{Timestamp: pb.TimeCreate}.ConvertToOptionalTime()
	p.DateTransaction = date.PbDate{Timestamp: pb.DateTransaction}.ConvertToDate()
	if p.Tags != nil {
		p.Tags = make([]model.Tag, len(pb.Tags))
		for i, tag := range pb.Tags {
			p.Tags[i] = PbTag{tag}.ConvertToStruct()
		}
	}
	return p
}

type PbTag struct {
	*pbTransaction.Tag
}

func (pb PbTag) ConvertToStruct() model.Tag {
	var p model.Tag
	p.TransactionID = pb.TransactionID
	p.TagID = pb.TagID
	return p
}

type PbCreateRes struct {
	*pbTransaction.CreateRes
}

func (pb PbCreateRes) ConvertToStruct() model.CreateRes {
	var p model.CreateRes
	p.ID = pb.ID
	return p
}
