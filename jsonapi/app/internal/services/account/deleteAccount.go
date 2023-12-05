package account

import (
	"context"
	"net/http"

	"jsonapi/app/internal/services/account/converter"
	"jsonapi/app/internal/services/account/model"
	"pkg/errors"
	"pkg/validation"

	"github.com/gorilla/schema"
)

// @Summary Удаление счета
// @Description Удаление данных по счету
// @Tags account
// @Security AuthJWT
// @Param Query query model.DeleteReq true "model.DeleteReq"
// @Produce jsonapi
// @Success 200 "Если удаление счета прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account [delete]
func (s *service) deleteAccount(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeDeleteAccountReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	if _, err := s.client.Delete(ctx, converter.DeleteReq{req}.ConvertToProto()); err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	return nil, nil
}

func decodeDeleteAccountReq(ctx context.Context, r *http.Request) (req model.DeleteReq, err error) {

	// Декодируем тело запроса в структуру
	if err = schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
