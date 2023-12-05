package converter

import (
	"auth/app/proto/pbAuth"

	"jsonapi/app/internal/services/auth/model"
)

type PbAuthRes struct {
	*pbAuth.AuthRes
}

func (pb PbAuthRes) ConvertToStruct() model.AuthRes {
	var p model.AuthRes
	p.Token = PbToken{pb.Token}.ConvertToStruct()
	p.ID = pb.ID
	return p
}

type PbRefreshTokensRes struct {
	*pbAuth.RefreshTokensRes
}

func (pb PbRefreshTokensRes) ConvertToStruct() model.RefreshTokensRes {
	var p model.RefreshTokensRes
	p.Token = PbToken{pb.Token}.ConvertToStruct()
	return p
}

type PbToken struct {
	*pbAuth.Token
}

func (pb PbToken) ConvertToStruct() model.Token {
	var p model.Token
	p.RefreshToken = pb.RefreshToken
	p.AccessToken = pb.AccessToken
	return p
}
