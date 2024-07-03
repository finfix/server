package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/transaction/model"
)

// @Summary Редактирование транзакции
// @Description Изменение данных транзакции
// @Tags transaction
// @Security AuthJWT
// @Accept json
// @Param Body body model.UpdateTransactionReq true "model.UpdateTransactionReq"
// @Success 200 "При успешном выполнении запроса возвращает пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /transaction [patch]
func (s *endpoint) updateTransaction(ctx context.Context, r *http.Request) (any, error) {

	var req model.UpdateTransactionReq

	// Декодируем запрос
	if err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, &req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.UpdateTransaction(ctx, req)
}
