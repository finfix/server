package model

type GetCurrenciesRes struct {
	Currencies []Currency `jsonapi:"-"`
}

type GetRes struct {
	Users []User `jsonapi:"-"`
}
