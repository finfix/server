package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/transaction/model"
)

// @Summary Удаление транзакции
// @Description Удаление данных транзакции и изменение баланса счетов
// @Tags transaction
// @Security AuthJWT
// @Param Query query model.DeleteTransactionReq true "model.DeleteTransactionReq"
// @Produce json
// @Success 200 "Если удаление транзакции прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,500 {object} errors.CustomError
// @Router /transaction [delete]
func (s *endpoint) deleteTransaction(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeSchema, model.DeleteTransactionReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.DeleteTransaction(ctx, req)
}
