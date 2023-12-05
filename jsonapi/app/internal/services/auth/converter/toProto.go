package converter

import (
	"auth/app/proto/pbAuth"

	"jsonapi/app/internal/services/auth/model"
)

type RefreshTokensReq struct {
	*model.RefreshTokensReq
}

func (p RefreshTokensReq) ConvertToProto() *pbAuth.RefreshTokensReq {
	pb := &pbAuth.RefreshTokensReq{}
	pb.Token = p.Token
	return pb
}

type SignInReq struct {
	*model.SignInReq
}

func (p SignInReq) ConvertToProto() *pbAuth.SignInReq {
	pb := &pbAuth.SignInReq{}
	pb.Email = p.Email
	pb.Password = p.Password
	pb.DeviceID = p.DeviceID
	return pb
}

type SignUpReq struct {
	*model.SignUpReq
}

func (p SignUpReq) ConvertToProto() *pbAuth.SignUpReq {
	pb := &pbAuth.SignUpReq{}
	pb.Email = p.Email
	pb.Password = p.Password
	pb.DeviceID = p.DeviceID
	pb.Name = p.Name
	return pb
}
