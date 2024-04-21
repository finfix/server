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

	// Декодируем параметры запроса в структуру
	req, err := decodeDeleteTagReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.DeleteTag(ctx, req)
}

func decodeDeleteTagReq(ctx context.Context, r *http.Request) (req model.DeleteTagReq, err error) {

	// Декодируем тело запроса в структуру
	if err := schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.Necessary, err = services.ExtractNecessaryFromCtx(ctx)
	if err != nil {
		return req, err
	}

	// Валидируем поля
	return req, validation.ZeroValue(req)
}
