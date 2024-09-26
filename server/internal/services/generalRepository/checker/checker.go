package checker

type CheckType string

const (
	Transactions  CheckType = "transactions"
	Accounts      CheckType = "accounts"
	AccountGroups CheckType = "account_groups"
	Tags          CheckType = "tags"
)
