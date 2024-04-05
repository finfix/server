package endpoint

import (
	"context"
	"net/http"

	"server/app/internal/services/account/model"
	"server/pkg/contextKeys"
	"server/pkg/validation"
)

// @Summary Получение списка групп счетов
// @Tags account
// @Security AuthJWT
// @Param Query query model.GetAccountGroupsReq true "model.GetAccountGroupsReq"
// @Produce json
// @Success 200 {object} []model.AccountGroup "[]model.AccountGroup"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account/accountGroups [get]
func (s *endpoint) getAccountGroups(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeGetAccountGroupsReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetAccountGroups(ctx, req)
}

func decodeGetAccountGroupsReq(ctx context.Context, _ *http.Request) (req model.GetAccountGroupsReq, err error) {

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value(contextKeys.UserIDKey).(uint32)
	req.DeviceID, _ = ctx.Value(contextKeys.DeviceIDKey).(string)

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
