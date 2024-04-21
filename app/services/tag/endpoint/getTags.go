package endpoint

import (
	"context"
	"net/http"

	"github.com/gorilla/schema"

	"server/app/pkg/errors"
	"server/app/pkg/validation"
	"server/app/services"
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

	// Декодируем параметры запроса в структуру
	req, err := decodeGetTagsReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetTags(ctx, req)
}

func decodeGetTagsReq(ctx context.Context, r *http.Request) (req model.GetTagsReq, err error) {

	// Декодируем параметры запроса в структуру
	if err = schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
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
