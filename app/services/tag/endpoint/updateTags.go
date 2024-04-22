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

// @Summary Редактирование транзакции
// @Description Изменение данных транзакции
// @Tags tag
// @Security AuthJWT
// @Accept json
// @Param Body body model.UpdateTagReq true "model.UpdateTagReq"
// @Success 200 "При успешном выполнении запроса возвращает пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /tag [patch]
func (s *endpoint) updateTag(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем параметры запроса в структуру
	req, err := decodeUpdateTagReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.UpdateTag(ctx, req)
}

func decodeUpdateTagReq(ctx context.Context, r *http.Request) (req model.UpdateTagReq, err error) {

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
