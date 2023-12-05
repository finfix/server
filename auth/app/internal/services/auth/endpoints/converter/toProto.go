package converter

import (
	"auth/app/internal/services/auth/model"
	"auth/app/proto/pbAuth"
)

type AuthRes struct {
	*model.AuthRes
}

func (p AuthRes) ConvertToProto() *pbAuth.AuthRes {
	pb := &pbAuth.AuthRes{}
	token := Token{&p.Token}
	pb.Token = token.ConvertToProto()
	pb.ID = p.ID
	return pb
}

type RefreshTokensRes struct {
	*model.RefreshTokensRes
}

func (p RefreshTokensRes) ConvertToProto() *pbAuth.RefreshTokensRes {
	pb := &pbAuth.RefreshTokensRes{}
	token := Token{&p.Token}
	pb.Token = token.ConvertToProto()
	return pb
}

type Token struct {
	*model.Token
}

func (p Token) ConvertToProto() *pbAuth.Token {

	pb := &pbAuth.Token{}
	pb.RefreshToken = p.RefreshToken
	pb.AccessToken = p.AccessToken
	return pb
}
