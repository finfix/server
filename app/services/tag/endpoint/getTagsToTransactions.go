package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/tag/model"
)

// @Summary Получение всех связей между подкатегориями и транзакциями
// @Tags tag
// @Security AuthJWT
// @Accept json
// @Param Body body model.GetTagsToTransactionsReq true "model.GetTagsToTransactionsReq"
// @Success 200 "При успешном выполнении запроса возвращает пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /tag/to_transactions [get]
func (s *endpoint) getTagsToTransaction(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeSchema, model.GetTagsToTransactionsReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetTagsToTransactions(ctx, req)
}
