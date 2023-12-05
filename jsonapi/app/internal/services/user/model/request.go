package model

type GetReq struct {
	ID uint32 `validate:"required" minimum:"1"`
}
