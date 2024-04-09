package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/validation"
	"server/app/services/account/model"
)

// @Summary Изменение порядковых мест двух счетов
// @Description Поменять два счета местами
// @Tags account
// @Security AuthJWT
// @Accept json
// @Param Body body model.SwitchReq true "model.SwitchReq"
// @Produce json
// @Success 200 "Если изменение порядка счетов прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account/switch [patch]
func (s *endpoint) switchAccounts(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeSwitchAccountsReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.Switch(ctx, req)
}

func decodeSwitchAccountsReq(ctx context.Context, r *http.Request) (req model.SwitchReq, err error) {

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
