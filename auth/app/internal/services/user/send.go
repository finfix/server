package user

import (
	"context"

	"auth/app/internal/services/user/converter"
	"auth/app/internal/services/user/model"
	"core/app/proto/pbUser"
)

type Service struct {
	user pbUser.UserClient
}

func New(user pbUser.UserClient) *Service {
	return &Service{
		user: user,
	}
}

func (c *Service) Get(ctx context.Context, req model.GetReq) (user []model.User, err error) {
	res, err := c.user.Get(ctx, converter.GetReq{&req}.ConvertToProto())
	if err != nil {
		return nil, err
	}
	return converter.PbGetRes{res}.ConvertToStruct().Users, nil
}

func (c *Service) Create(ctx context.Context, userData model.CreateReq) (uint32, error) {
	res, err := c.user.Create(ctx, converter.CreateReq{&userData}.ConvertToProto())
	if err != nil {
		return 0, err
	}
	return converter.PbCreateRes{res}.ConvertToStruct().ID, nil
}
