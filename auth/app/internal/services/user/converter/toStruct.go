package converter

import (
	"auth/app/internal/services/user/model"
	"core/app/proto/pbUser"
	"pkg/datetime/time"
)

type PbGetRes struct {
	*pbUser.GetRes
}

func (p PbGetRes) ConvertToStruct() model.GetRes {
	var res model.GetRes
	res.Users = make([]model.User, len(p.Users))
	for i, user := range p.Users {
		res.Users[i] = PbUser{User: user}.ConvertToStruct()
	}
	return res
}

type PbUser struct {
	*pbUser.User
}

func (pb PbUser) ConvertToStruct() model.User {
	var p model.User
	p.ID = pb.ID
	p.Name = pb.Name
	p.Email = pb.Email
	p.PasswordHash = pb.PasswordHash
	p.VerificationEmailCode = pb.VerificationEmailCode
	p.TimeCreate = time.PbTime{pb.TimeCreate}.ConvertToTime()
	return p
}

type PbCreateRes struct {
	*pbUser.CreateRes
}

func (p PbCreateRes) ConvertToStruct() model.CreateRes {
	var res model.CreateRes
	res.ID = p.ID
	return res
}
