package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/validation"
	"server/app/services"
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

	// Декодируем параметры запроса в структуру
	req, err := decodeGetTagsToTransactionReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetTagsToTransactions(ctx, req)
}

func decodeGetTagsToTransactionReq(ctx context.Context, _ *http.Request) (req model.GetTagsToTransactionsReq, err error) {

	// Заполняем поля из контекста
	req.Necessary, err = services.ExtractNecessaryFromCtx(ctx)
	if err != nil {
		return req, err
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
