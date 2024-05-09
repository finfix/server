package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/account/model"
)

// @Summary Изменение порядковых мест двух счетов
// @Description Поменять два счета местами
// @Tags account
// @Security AuthJWT
// @Accept json
// @Param Body body model.SwitchAccountBetweenThemselvesReq true "model.SwitchAccountBetweenThemselvesReq"
// @Produce json
// @Success 200 "Если изменение порядка счетов прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account/switch [patch]
func (s *endpoint) switchAccounts(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, model.SwitchAccountBetweenThemselvesReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.SwitchAccountBetweenThemselves(ctx, req)
}
