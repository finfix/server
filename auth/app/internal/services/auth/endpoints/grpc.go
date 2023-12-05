package endpoints

import (
	"context"

	"logger/app/logging"

	"auth/app/internal/services/auth/endpoints/converter"
	"auth/app/internal/services/auth/model"
	pb "auth/app/proto/pbAuth"
)

type AuthService interface {
	SignIn(context.Context, model.SignInReq) (model.AuthRes, error)
	SignUp(context.Context, model.SignUpReq) (model.AuthRes, error)
	RefreshTokens(context.Context, string) (model.RefreshTokensRes, error)
}

func (s *Endpoint) SignIn(ctx context.Context, in *pb.SignInReq) (*pb.AuthRes, error) {

	res, err := s.service.SignIn(ctx, converter.PbSignInReq{in}.ConvertToStruct())
	if err != nil {
		return nil, err
	}

	return converter.AuthRes{&res}.ConvertToProto(), nil
}

func (s *Endpoint) SignUp(ctx context.Context, in *pb.SignUpReq) (*pb.AuthRes, error) {

	res, err := s.service.SignUp(ctx, converter.PbSignUpReq{in}.ConvertToStruct())
	if err != nil {
		return nil, err
	}

	return converter.AuthRes{&res}.ConvertToProto(), nil
}

func (s *Endpoint) RefreshTokens(ctx context.Context, in *pb.RefreshTokensReq) (*pb.RefreshTokensRes, error) {

	res, err := s.service.RefreshTokens(ctx, converter.PbRefreshTokensReq{in}.ConvertToStruct().Token)
	if err != nil {
		return nil, err
	}

	return converter.RefreshTokensRes{&res}.ConvertToProto(), nil
}

type Endpoint struct {
	pb.UnsafeAuthServer
	service AuthService
	logger  *logging.Logger
}

func New(service AuthService, logger *logging.Logger) pb.AuthServer {
	return &Endpoint{
		service: service,
		logger:  logger,
	}
}
