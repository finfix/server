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

// @Summary Удаление счета
// @Description Удаление данных по счету
// @Tags account
// @Security AuthJWT
// @Param Query query model.DeleteReq true "model.DeleteReq"
// @Produce json
// @Success 200 "Если удаление счета прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account [delete]
func (s *endpoint) deleteAccount(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeDeleteAccountReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.Delete(ctx, req)
}

func decodeDeleteAccountReq(ctx context.Context, r *http.Request) (req model.DeleteReq, err error) {

	// Декодируем тело запроса в структуру
	if err = schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value(contextKeys.UserIDKey).(uint32)
	req.DeviceID, _ = ctx.Value(contextKeys.DeviceIDKey).(string)

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
