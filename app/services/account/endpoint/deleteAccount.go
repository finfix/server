package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/account/model"
)

// @Summary Удаление счета
// @Description Удаление данных по счету
// @Tags account
// @Security AuthJWT
// @Param Query query model.DeleteAccountReq true "model.DeleteAccountReq"
// @Produce json
// @Success 200 "Если удаление счета прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account [delete]
func (s *endpoint) deleteAccount(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeSchema, model.DeleteAccountReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.DeleteAccount(ctx, req)
}
