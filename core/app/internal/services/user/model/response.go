package model

type GetRes struct {
	Users []User
}

type CreateRes struct {
	ID uint32
}

type GetCurrenciesRes struct {
	Currencies []Currency
}
