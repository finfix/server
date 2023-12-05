package account

import (
	"context"
	"encoding/json"
	"net/http"

	"jsonapi/app/internal/services/account/converter"
	"jsonapi/app/internal/services/account/model"
	"pkg/errors"
	"pkg/validation"
)

// @Summary Редактирование счета
// @Description Изменение данных счета
// @Tags account
// @Security AuthJWT
// @Accept jsonapi
// @Param Body body model.UpdateReq true "model.UpdateReq"
// @Produce jsonapi
// @Success 200 "Если редактирование счета прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account [patch]
func (s *service) updateAccount(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeUpdateAccountReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	if _, err = s.client.Update(ctx, converter.UpdateReq{req}.ConvertToProto()); err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	return nil, nil
}

func decodeUpdateAccountReq(ctx context.Context, r *http.Request) (req model.UpdateReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
