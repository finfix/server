package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/accountGroup/model"
)

// @Summary Удаление группы счетов
// @Tags accountGroup
// @Security AuthJWT
// @Param Query query model.DeleteAccountGroupReq true "model.DeleteAccountGroupReq"
// @Produce json
// @Success 200 "Если удаление группы счетов прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /accountGroup [delete]
func (s *endpoint) deleteAccountGroup(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeSchema, model.DeleteAccountGroupReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.DeleteAccountGroup(ctx, req)
}
