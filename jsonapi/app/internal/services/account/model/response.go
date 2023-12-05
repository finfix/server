package model

type GetRes struct {
	Accounts []Account
}

type CreateRes struct {
	ID uint32 `json:"id"` // Идентификатор созданного счета
}

type QuickStatisticRes struct {
	QuickStatistic []QuickStatistic
}

type GetAccountGroupsRes struct {
	AccountGroups []AccountGroup
}
