package endpoint

import (
	"context"
	"net/http"

	"pkg/http/decoder"

	"server/internal/services/transaction/model"
)

// @Summary Создание транзакции
// @Description Создание транзакции и изменение баланса счетов, между которыми она произошла
// @Tags transaction
// @Security AuthJWT
// @Accept json
// @Param Body body model.CreateTransactionReq true "model.CreateTransactionReq"
// @Produce json
// @Success 200 {object} model.CreateTransactionRes
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /transaction [post]
func (s *endpoint) createTransaction(ctx context.Context, r *http.Request) (any, error) {

	var req model.CreateTransactionReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeJSON); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	id, err := s.service.CreateTransaction(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.CreateTransactionRes{ID: id}, nil
}
