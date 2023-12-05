package account

import (
	"context"
	"net/http"

	"jsonapi/app/internal/services/account/converter"
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

	// Вызываем метод сервиса
	groups, err := s.client.GetAccountGroups(ctx, converter.GetAccountGroupsReq{req}.ConvertToProto())
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	return converter.PbGetAccountGroupsRes{groups}.ConvertToStruct().AccountGroups, nil
}

func decodeGetAccountGroupsReq(ctx context.Context, _ *http.Request) (req model.GetAccountGroupsReq, err error) {

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
