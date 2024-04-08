package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/services/account/model"
	"server/pkg/contextKeys"
	"server/pkg/errors"
	"server/pkg/validation"
)

// @Summary Редактирование счета
// @Description Изменение данных счета
// @Tags account
// @Security AuthJWT
// @Accept json
// @Param Body body model.UpdateReq true "model.UpdateReq"
// @Produce json
// @Success 200 "Если редактирование счета прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account [patch]
func (s *endpoint) updateAccount(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeUpdateAccountReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.Update(ctx, req)
}

func decodeUpdateAccountReq(ctx context.Context, r *http.Request) (req model.UpdateReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value(contextKeys.UserIDKey).(uint32)
	req.DeviceID, _ = ctx.Value(contextKeys.DeviceIDKey).(string)

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
