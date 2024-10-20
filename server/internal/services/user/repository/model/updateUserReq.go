package model

type UpdateUserReq struct {
	ID              uint32
	Name            *string
	Email           *string
	PasswordHash    *[]byte
	PasswordSalt    *[]byte
	DefaultCurrency *string
}
