package service

import (
	"context"
	"encoding/json"
	"logger/app/logging"
	"net/http"
	"net/url"
	"pkg/errors"
	"time"

	"core/app/internal/config"
)

type Repository interface {
	UpdCurrencies(ctx context.Context, rates map[string]float64) error
}

// UpdateCurrencies обновляет курсы валют
func (s *Service) UpdateCurrencies(ctx context.Context) (map[string]float64, error) {

	var providerModel struct {
		Meta struct {
			LastUpdatedAt time.Time `json:"last_updated_at"`
		} `json:"meta"`
		Rates map[string]struct {
			Rate float64 `json:"value"`
		} `json:"data"`
	}

	urlValues := url.Values{}

	// URL для получения курсов валют
	urlString := "https://api.currencyapi.com/v3/latest"

	// Параметры запроса
	urlValues.Add("apikey", config.GetConfig().ApiKeys.CurrencyProvider)

	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	u.RawQuery = urlValues.Encode()

	// Отправляем запрос
	resp, err := http.Get(u.String())
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
		return nil, errors.BadGateway.NewCtx("Error while getting currency rates", "HTTP code: %v", resp.StatusCode)
	}

	rates := make(map[string]float64, len(providerModel.Rates))

	// Конвертируем полученные данные в нужный формат
	for currency, rate := range providerModel.Rates {
		rates[currency] = rate.Rate
	}

	// Обновляем курсы валют в БД
	return rates, s.repository.UpdCurrencies(ctx, rates)
}

type Service struct {
	repository Repository
	logger     *logging.Logger
}

func New(rep Repository, logger *logging.Logger) *Service {
	return &Service{
		repository: rep,
		logger:     logger,
	}
}
