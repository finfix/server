package model

type UserToAccountGroup struct {
	UserID         uint32 `db:"user_id"`
	AccountGroupID uint32 `db:"account_group_id"`
}
