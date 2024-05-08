package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/tag/model"
)

// @Summary Удаление транзакции
// @Description Удаление данных транзакции и изменение баланса счетов
// @Tags tag
// @Security AuthJWT
// @Param Query query model.DeleteTagReq true "model.DeleteTagReq"
// @Produce json
// @Success 200 "Если удаление транзакции прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,500 {object} errors.CustomError
// @Router /tag [delete]
func (s *endpoint) deleteTag(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeSchema, model.DeleteTagReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.DeleteTag(ctx, req)
}
