package converter

import (
	"core/app/proto/pbUser"
	"jsonapi/app/internal/services/user/model"
)

type GetReq struct {
	model.GetReq
}

func (p GetReq) ConvertToProto() *pbUser.GetReq {
	var pb pbUser.GetReq
	pb.IDs = []uint32{p.ID}
	return &pb
}
