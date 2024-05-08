package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/transaction/model"
)

// @Summary Получение всех транзакций
// @Description Получение всех транзакций по фильтрам
// @Tags transaction
// @Security AuthJWT
// @Param Query query model.GetTransactionsReq true "model.CreateTransactionReq"
// @Produce json
// @Success 200 {object} []model.Transaction
// @Failure 400,404,500 {object} errors.CustomError
// @Router /transaction [get]
func (s *endpoint) getTransactions(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeSchema, model.GetTransactionsReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetTransactions(ctx, req)
}
