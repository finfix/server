package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/pkg/errors"
	"server/app/pkg/validation"
	"server/app/services"
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

	// Декодируем тело запроса в структуру
	req, err := decodeCreateTagReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	id, err := s.service.CreateTag(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.CreateTagRes{ID: id}, nil
}

func decodeCreateTagReq(ctx context.Context, r *http.Request) (req model.CreateTagReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.Necessary, err = services.ExtractNecessaryFromCtx(ctx)
	if err != nil {
		return req, err
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
