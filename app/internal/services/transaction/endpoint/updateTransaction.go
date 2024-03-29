package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/internal/services/transaction/model"
	"server/pkg/errors"
	"server/pkg/validation"
)

// @Summary Редактирование транзакции
// @Description Изменение данных транзакции
// @Tags transaction
// @Security AuthJWT
// @Accept json
// @Param Body body model.UpdateReq true "model.CreateReq"
// @Produce json
// @Success 200 "Если редактирование транзакции прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /transaction [patch]
func (s *endpoint) updateTransaction(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем параметры запроса в структуру
	req, err := decodeUpdateTransactionReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.Update(ctx, req)
}

func decodeUpdateTransactionReq(ctx context.Context, r *http.Request) (req model.UpdateReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Валидируем поля
	if req.AmountFrom != nil && *req.AmountFrom <= 0 {
		return req, errors.BadRequest.New("amountFrom must be greater than 0")
	}
	if req.AmountTo != nil && *req.AmountTo <= 0 {
		return req, errors.BadRequest.New("amountTo must be greater than 0")
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
