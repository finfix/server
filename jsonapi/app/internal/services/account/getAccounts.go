package account

import (
	"context"
	pb "core/app/proto/pbAccount"
	"net/http"
	"pkg/converter"

	"jsonapi/app/internal/services/account/model"
	"pkg/errors"
	"pkg/validation"

	"github.com/gorilla/schema"
)

// @Summary Получение счетов по фильтрам
// @Tags account
// @Security AuthJWT
// @Param Query query model.GetReq false "model.GetReq"
// @Produce json
// @Success 200 {object} []model.Account
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account [get]
func (s *service) get(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeGetReq(ctx, r)
	if err != nil {
		return nil, err
	}

	in, err := converter.Convert(pb.GetReq{}, req)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	out, err := s.client.Get(ctx, &in)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	res, err := converter.Convert(model.GetRes{}, out)
	if err != nil {
		return nil, err
	}

	// Конвертируем ответ во внутреннюю структуру
	return res.Accounts, nil
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
