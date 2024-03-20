package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/internal/services/account/model"
	"server/pkg/errors"
	"server/pkg/validation"
)

// @Summary Создание счета
// @Description Создается новый счет, если у него есть остаток, то создается транзакция от нулевого счета для баланса
// @Tags account
// @Security AuthJWT
// @Accept json
// @Param Body body model.CreateReq true "model.CreateReq"
// @Produce json
// @Success 200 {object} model.CreateRes
// @Failure 400,401,403,500 {object} errors.CustomError
// @Router /account [post]
func (s *endpoint) createAccount(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeCreateAccountReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.Create(ctx, req)
}

func decodeCreateAccountReq(ctx context.Context, r *http.Request) (req model.CreateReq, err error) {

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

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
