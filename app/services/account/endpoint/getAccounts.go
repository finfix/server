package endpoint

import (
	"context"
	"net/http"

	"github.com/gorilla/schema"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/validation"
	"server/app/services/account/model"
)

// @Summary Получение счетов по фильтрам
// @Tags account
// @Security AuthJWT
// @Param Query query model.GetReq false "model.GetReq"
// @Produce json
// @Success 200 {object} []model.Account
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account [get]
func (s *endpoint) get(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeGetReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.Get(ctx, req)
}

func decodeGetReq(ctx context.Context, r *http.Request) (req model.GetReq, err error) {

	// Декодируем параметры запроса в структуру
	if err = schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value(contextKeys.UserIDKey).(uint32)
	req.DeviceID, _ = ctx.Value(contextKeys.DeviceIDKey).(string)

	// Валидируем поля
	if err = req.Type.Validate(); err != nil {
		return req, err
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
