package endpoint

import (
	"context"
	"net/http"

	"github.com/gorilla/schema"

	"server/app/pkg/errors"
	"server/app/pkg/validation"
	"server/app/services"
	"server/app/services/transaction/model"
)

// @Summary Получение всех транзакций
// @Description Получение всех транзакций по фильтрам
// @Tags transaction
// @Security AuthJWT
// @Param Query query model.GetTransactionsReq true "model.CreateTransactionReq"
// @Produce json
// @Success 200 {object} []model.Transaction
// @Failure 400,404,500 {object} errors.CustomError
// @Router /transaction [get]
func (s *endpoint) getTransactions(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем параметры запроса в структуру
	req, err := decodeGetTransactionsReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetTransactions(ctx, req)
}

func decodeGetTransactionsReq(ctx context.Context, r *http.Request) (req model.GetTransactionsReq, err error) {

	// Декодируем параметры запроса в структуру
	if err = schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.Necessary, err = services.ExtractNecessaryFromCtx(ctx)
	if err != nil {
		return req, err
	}

	// Валидируем поля
	if err = req.Type.Validate(); err != nil {
		return req, err
	}
	if req.DateFrom != nil && req.DateTo != nil {
		if req.DateFrom.After(req.DateTo.Time) || req.DateFrom.Equal(req.DateTo.Time) {
			return req, errors.BadRequest.New("date_from must be less than date_to")
		}
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}