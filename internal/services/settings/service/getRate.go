package service

import "github.com/shopspring/decimal"

func getRate(rates map[string]decimal.Decimal, currency, currencyRelate string) decimal.Decimal {
	currencyRate := rates[currency]
	currencyRelateRate := rates[currencyRelate]
	if currencyRate.GreaterThan(currencyRelateRate) {
		return currencyRate.Div(currencyRelateRate)

	}
	return currencyRelateRate.Div(currencyRate)
}
