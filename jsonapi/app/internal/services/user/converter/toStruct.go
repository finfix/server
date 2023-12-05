package converter

import (
	"core/app/proto/pbUser"
	"jsonapi/app/internal/services/user/model"
	"pkg/datetime/time"
)

type PbGetCurrenciesRes struct {
	*pbUser.GetCurrenciesRes
}

func (pb PbGetCurrenciesRes) ConvertToStruct() model.GetCurrenciesRes {
	var p model.GetCurrenciesRes
	p.Currencies = make([]model.Currency, len(pb.Currencies))
	for i, currency := range pb.Currencies {
		p.Currencies[i] = PbCurrency{currency}.ConvertToStruct()
	}
	return p
}

type PbCurrency struct {
	*pbUser.Currency
}

func (pb PbCurrency) ConvertToStruct() model.Currency {
	var p model.Currency
	p.Signatura = pb.Signatura
	p.Symbol = pb.Symbol
	p.Name = pb.Name
	p.Rate = pb.Rate
	return p
}

type PbGetRes struct {
	*pbUser.GetRes
}

func (pb PbGetRes) ConvertToStruct() model.GetRes {
	var p model.GetRes
	p.Users = make([]model.User, len(pb.Users))
	for i, user := range pb.Users {
		p.Users[i] = PbUser{user}.ConvertToStruct()
	}
	return p
}

type PbUser struct {
	*pbUser.User
}

func (pb PbUser) ConvertToStruct() model.User {
	var p model.User
	p.ID = pb.ID
	p.Email = pb.Email
	p.Name = pb.Name
	p.TimeCreate = time.PbTime{pb.TimeCreate}.ConvertToOptionalTime()
	p.DefaultCurrency = pb.DefaultCurrency
	return p
}
