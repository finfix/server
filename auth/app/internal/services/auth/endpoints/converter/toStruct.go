package converter

import (
	"auth/app/internal/services/auth/model"
	"auth/app/proto/pbAuth"
)

type PbRefreshTokensReq struct {
	*pbAuth.RefreshTokensReq
}

func (pb PbRefreshTokensReq) ConvertToStruct() model.RefreshTokensReq {
	var p model.RefreshTokensReq
	p.Token = pb.Token
	return p
}

type PbSignInReq struct {
	*pbAuth.SignInReq
}

func (pb PbSignInReq) ConvertToStruct() model.SignInReq {
	var p model.SignInReq
	p.Email = pb.Email
	p.Password = pb.Password
	p.DeviceID = pb.DeviceID
	return p
}

type PbSignUpReq struct {
	*pbAuth.SignUpReq
}

func (pb PbSignUpReq) ConvertToStruct() model.SignUpReq {
	var p model.SignUpReq
	p.Email = pb.Email
	p.Password = pb.Password
	p.DeviceID = pb.DeviceID
	p.Name = pb.Name
	return p
}
