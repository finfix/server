package model

type UpdateCurrenciesRes struct {
	Rates map[string]float64 `json:"rates"`
}
