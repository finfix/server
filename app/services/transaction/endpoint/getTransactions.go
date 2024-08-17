package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/http/decoder"
	"server/app/services/transaction/model"
)

// @Summary Получение всех транзакций
// @Description Получение всех транзакций по фильтрам
// @Tags transaction
// @Security AuthJWT
// @Param Query query model.GetTransactionsReq true "model.CreateTransactionReq"
// @Produce json
// @Success 200 {object} []model.Transaction
// @Failure 400,404,500 {object} errors.Error
// @Router /transaction [get]
func (s *endpoint) getTransactions(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetTransactionsReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetTransactions(ctx, req)
}
