package endpoint

import (
	"context"
	"encoding/json"
	"github.com/shopspring/decimal"
	"net/http"

	"server/app/pkg/errors"
	"server/app/pkg/validation"
	"server/app/services"
	"server/app/services/transaction/model"
)

// @Summary Создание транзакции
// @Description Создание транзакции и изменение баланса счетов, между которыми она произошла
// @Tags transaction
// @Security AuthJWT
// @Accept json
// @Param Body body model.CreateTransactionReq true "model.CreateTransactionReq"
// @Produce json
// @Success 200 {object} model.CreateTransactionRes
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /transaction [post]
func (s *endpoint) createTransaction(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем тело запроса в структуру
	req, err := decodeCreateTransactionReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	id, err := s.service.CreateTransaction(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.CreateTransactionRes{ID: id}, nil
}

func decodeCreateTransactionReq(ctx context.Context, r *http.Request) (req model.CreateTransactionReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
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
	if req.AmountFrom.LessThanOrEqual(decimal.Zero) || req.AmountTo.LessThanOrEqual(decimal.Zero) {
		return req, errors.BadRequest.New("amountFrom and amountTo must be greater than 0")
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
