package endpoint

import (
	"context"
	"net/http"

	"github.com/gorilla/schema"

	"server/app/internal/services/transaction/model"
	"server/pkg/errors"
	"server/pkg/validation"
)

// @Summary Удаление транзакции
// @Description Удаление данных транзакции и изменение баланса счетов
// @Tags transaction
// @Security AuthJWT
// @Param Query query model.DeleteReq true "model.DeleteReq"
// @Produce json
// @Success 200 "Если удаление транзакции прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,500 {object} errors.CustomError
// @Router /transaction [delete]
func (s *endpoint) deleteTransaction(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем параметры запроса в структуру
	req, err := decodeDeleteTransactionReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.Delete(ctx, req)
}

func decodeDeleteTransactionReq(ctx context.Context, r *http.Request) (req model.DeleteReq, err error) {

	// Декодируем тело запроса в структуру
	if err := schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Валидируем поля
	return req, validation.ZeroValue(req)
}
