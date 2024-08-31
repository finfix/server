package endpoint

import (
	"context"
	"net/http"

	"server/pkg/http/decoder"
	"server/internal/services/auth/model"
)

// @Summary Выход пользователя из приложения
// @Tags auth
// @Accept json
// @Produce json
// @Security AuthJWT
// @Success 200 "При успешном выходе возвращается null"
// @Failure 400,404,500 {object} errors.Error
// @Router /auth/signOut [post]
func (s *endpoint) signOut(ctx context.Context, r *http.Request) (any, error) {

	var req model.SignOutReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeJSON); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.SignOut(ctx, req)
}
