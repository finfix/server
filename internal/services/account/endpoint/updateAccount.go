package endpoint

import (
	"context"
	"net/http"

	"server/pkg/http/decoder"
	"server/internal/services/account/model"
)

// @Summary Редактирование счета
// @Description Изменение данных счета
// @Tags account
// @Security AuthJWT
// @Accept json
// @Param Body body model.UpdateAccountReq true "model.UpdateAccountReq"
// @Produce json
// @Success 200 "Если редактирование счета прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /account [patch]
func (s *endpoint) updateAccount(ctx context.Context, r *http.Request) (any, error) {

	var req model.UpdateAccountReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeJSON); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.UpdateAccount(ctx, req)
}
