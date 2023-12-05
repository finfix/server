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

// @Summary Получение счетов по фильтрам
// @Tags account
// @Security AuthJWT
// @Param Query query model.GetReq false "model.GetReq"
// @Produce jsonapi
// @Success 200 {object} []model.Account
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account [get]
func (s *service) get(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeGetReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	proto, err := s.client.Get(ctx, converter.GetReq{req}.ConvertToProto())
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Конвертируем ответ во внутреннюю структуру
	return converter.PbGetRes{proto}.ConvertToStruct().Accounts, nil
}

func decodeGetReq(ctx context.Context, r *http.Request) (req model.GetReq, err error) {

	// Декодируем параметры запроса в структуру
	if err = schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
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
