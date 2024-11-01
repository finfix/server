package network

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/shopspring/decimal"

	"pkg/errors"
)

func GetCurrencyRates(ctx context.Context, apiKey string) (map[string]decimal.Decimal, error) {
	var providerModel struct {
		Meta struct {
			LastUpdatedAt time.Time `json:"last_updated_at"`
		} `json:"meta"`
		Rates map[string]struct {
			Rate decimal.Decimal `json:"value"`
		} `json:"data"`
	}

	urlValues := url.Values{}

	// URL для получения курсов валют
	urlString := "https://api.currencyapi.com/v3/latest"

	// Параметры запроса
	urlValues.Add("apikey", apiKey)

	uri, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	uri.RawQuery = urlValues.Encode()

	// Отправляем запрос
	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	req = req.WithContext(ctx)
	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.BadGateway.Wrap(err)
	}
	defer resp.Body.Close()

	// Смотрим код ответа
	switch resp.StatusCode {
	case http.StatusOK:

		// Декодируем ответ
		if err = json.NewDecoder(resp.Body).Decode(&providerModel); err != nil {
			return nil, errors.InternalServer.Wrap(err)
		}
	default:
		return nil, errors.BadGateway.New("Error while getting currency rates",
			errors.ParamsOption("HTTP code", resp.StatusCode),
		)
	}

	rates := make(map[string]decimal.Decimal, len(providerModel.Rates))

	// Конвертируем полученные данные в нужный формат
	for currency, rate := range providerModel.Rates {
		rates[currency] = rate.Rate
	}

	return rates, nil
}
