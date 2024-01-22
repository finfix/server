package account

import (
	"context"
	pb "core/app/proto/pbAccount"
	"net/http"
	"pkg/converter"

	"jsonapi/app/internal/services/account/model"
	"pkg/errors"
	"pkg/validation"
)

// @Summary Получение списка групп счетов
// @Tags account
// @Security AuthJWT
// @Param Query query model.GetAccountGroupsReq true "model.GetAccountGroupsReq"
// @Produce json
// @Success 200 {object} []model.AccountGroup "[]model.AccountGroup"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account/accountGroups [get]
func (s *service) getAccountGroups(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeGetAccountGroupsReq(ctx, r)
	if err != nil {
		return nil, err
	}

	in, err := converter.Convert(pb.GetAccountGroupsReq{}, req)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	out, err := s.client.GetAccountGroups(ctx, &in)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	res, err := converter.Convert(model.GetAccountGroupsRes{}, out)
	if err != nil {
		return nil, err
	}

	// Конвертируем ответ во внутреннюю структуру
	return res, nil
}

func decodeGetAccountGroupsReq(ctx context.Context, _ *http.Request) (req model.GetAccountGroupsReq, err error) {

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
