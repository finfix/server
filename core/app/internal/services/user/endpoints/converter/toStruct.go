package converter

import (
	"core/app/internal/services/user/model"
	"core/app/proto/pbUser"
	"pkg/datetime/time"
)

type PbGetReq struct {
	*pbUser.GetReq
}

func (pb PbGetReq) ConvertToStruct() model.GetReq {
	var p model.GetReq
	p.IDs = pb.IDs
	p.Emails = pb.Emails
	return p
}

type PbCreateReq struct {
	*pbUser.CreateReq
}

func (pb PbCreateReq) ConvertToStruct() model.CreateReq {
	var p model.CreateReq
	p.Name = pb.Name
	p.Email = pb.Email
	p.PasswordHash = pb.PasswordHash
	p.DefaultCurrency = pb.DefaultCurrency
	p.TimeCreate = time.PbTime{pb.TimeCreate}.ConvertToTime()
	return p
}
