package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/tag/model"
)

// @Summary Создание подкатегории
// @Description Создание подкатегории
// @Tags tag
// @Security AuthJWT
// @Accept json
// @Param Body body model.CreateTagReq true "model.CreateTagReq"
// @Produce json
// @Success 200 {object} model.CreateTagRes
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /tag [post]
func (s *endpoint) createTag(ctx context.Context, r *http.Request) (any, error) {

	var req model.CreateTagReq

	// Декодируем запрос
	if err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, &req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	id, err := s.service.CreateTag(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.CreateTagRes{ID: id}, nil
}
