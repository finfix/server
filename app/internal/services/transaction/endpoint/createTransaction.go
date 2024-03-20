package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/internal/services/transaction/model"
	"server/pkg/errors"
	"server/pkg/validation"
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
	return s.service.Create(ctx, req)
}

func decodeCreateTransactionReq(ctx context.Context, r *http.Request) (req model.CreateReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

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
