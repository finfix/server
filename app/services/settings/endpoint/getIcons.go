package endpoint

import (
	"context"
	"net/http"

	_ "server/app/services/user/model" //nolint:golint
)

// @Summary Получение иконок с их адресами
// @Tags settings
// @Security AuthJWT
// @Produce json
// @Success 200 {object} []model.Icon
// @Failure 401,500 {object} errors.Error
// @Router /settings/icons [get]
func (s *endpoint) getIcons(ctx context.Context, _ *http.Request) (any, error) {

	// Вызываем метод сервиса
	return s.service.GetIcons(ctx)
}
