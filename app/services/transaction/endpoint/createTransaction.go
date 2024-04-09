package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/validation"
	model2 "server/app/services/transaction/model"
)

// @Summary Создание транзакции
// @Description Создание транзакции и изменение баланса счетов, между которыми она произошла
// @Tags transaction
// @Security AuthJWT
// @Accept json
// @Param Body body model.CreateReq true "model.CreateReq"
// @Produce json
// @Success 200 {object} model.CreateRes
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /transaction [post]
func (s *endpoint) createTransaction(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем тело запроса в структуру
	req, err := decodeCreateTransactionReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	id, err := s.service.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return model2.CreateRes{ID: id}, nil
}

func decodeCreateTransactionReq(ctx context.Context, r *http.Request) (req model2.CreateReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value(contextKeys.UserIDKey).(uint32)
	req.DeviceID, _ = ctx.Value(contextKeys.DeviceIDKey).(string)

	// Валидируем поля
	if err = req.Type.Validate(); err != nil {
		return req, err
	}
	if req.AmountFrom <= 0 || req.AmountTo <= 0 {
		return req, errors.BadRequest.New("amountFrom and amountTo must be greater than 0")
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
