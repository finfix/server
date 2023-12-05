package converter

import (
	"core/app/internal/services/user/model"
	"core/app/proto/pbUser"
	"pkg/datetime/time"
)

type GetRes struct {
	*model.GetRes
}

func (p GetRes) ConvertToProto() *pbUser.GetRes {
	var pb pbUser.GetRes
	pb.Users = make([]*pbUser.User, len(p.Users))
	for i, user := range p.Users {
		pb.Users[i] = User{&user}.ConvertToProto()
	}
	return &pb
}

type User struct {
	*model.User
}

func (p User) ConvertToProto() *pbUser.User {
	var pb pbUser.User
	pb.ID = p.ID
	pb.Name = p.Name
	pb.Email = p.Email
	pb.PasswordHash = p.PasswordHash
	pb.VerificationEmailCode = p.VerificationEmailCode
	pb.FCMToken = p.FCMToken
	pb.DefaultCurrency = p.DefaultCurrency
	pb.TimeCreate = time.Time{p.TimeCreate}.ConvertToProto()
	return &pb
}

type CreateRes struct {
	*model.CreateRes
}

func (p CreateRes) ConvertToProto() *pbUser.CreateRes {
	var pb pbUser.CreateRes
	pb.ID = p.ID
	return &pb
}

type Currency struct {
	*model.Currency
}

func (p Currency) ConvertToProto() *pbUser.Currency {
	var pb pbUser.Currency
	pb.Signatura = p.Signatura
	pb.Name = p.Name
	pb.Symbol = p.Symbol
	pb.Rate = p.Rate
	return &pb
}

type GetCurrenciesRes struct {
	*model.GetCurrenciesRes
}

func (p GetCurrenciesRes) ConvertToProto() *pbUser.GetCurrenciesRes {
	var pb pbUser.GetCurrenciesRes
	pb.Currencies = make([]*pbUser.Currency, len(p.Currencies))
	for i, currency := range p.Currencies {
		pb.Currencies[i] = Currency{&currency}.ConvertToProto()
	}
	return &pb
}
