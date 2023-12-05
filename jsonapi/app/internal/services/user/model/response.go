package model

type GetCurrenciesRes struct {
	Currencies []Currency `json:"-"`
}

type GetRes struct {
	Users []User `json:"-"`
}
