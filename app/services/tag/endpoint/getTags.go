package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/tag/model"
)

// @Summary Получение всех транзакций
// @Description Получение всех транзакций по фильтрам
// @Tags tag
// @Security AuthJWT
// @Param Query query model.GetTagsReq true "model.CreateTagReq"
// @Produce json
// @Success 200 {object} []model.Tag
// @Failure 400,404,500 {object} errors.CustomError
// @Router /tag [get]
func (s *endpoint) getTags(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetTagsReq

	// Декодируем запрос
	if err := middleware.DefaultDecoder(ctx, r, middleware.DecodeSchema, &req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetTags(ctx, req)
}
