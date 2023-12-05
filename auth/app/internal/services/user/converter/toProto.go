package converter

import (
	"auth/app/internal/services/user/model"
	"core/app/proto/pbUser"
	"pkg/datetime/time"
)

type GetReq struct {
	*model.GetReq
}

func (p GetReq) ConvertToProto() *pbUser.GetReq {
	var pb pbUser.GetReq
	pb.IDs = p.IDs
	pb.Emails = p.Emails
	return &pb
}

type CreateReq struct {
	*model.CreateReq
}

func (p CreateReq) ConvertToProto() *pbUser.CreateReq {
	var pb pbUser.CreateReq
	pb.Name = p.Name
	pb.Email = p.Email
	pb.PasswordHash = p.PasswordHash
	pb.DefaultCurrency = p.DefaultCurrency
	pb.TimeCreate = time.Time{p.TimeCreate}.ConvertToProto()
	return &pb
}
