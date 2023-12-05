package transaction

import (
	"context"
	"net/http"

	"jsonapi/app/internal/services/transaction/converter"
	"jsonapi/app/internal/services/transaction/model"
	"pkg/errors"
	"pkg/validation"

	"github.com/gorilla/schema"
)

// @Summary Удаление транзакции
// @Description Удаление данных транзакции и изменение баланса счетов
// @Tags transaction
// @Security AuthJWT
// @Param Query query model.DeleteReq true "model.DeleteReq"
// @Produce jsonapi
// @Success 200 "Если удаление транзакции прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,500 {object} errors.CustomError
// @Router /transaction [delete]
func (s *service) deleteTransaction(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем параметры запроса в структуру
	req, err := decodeDeleteTransactionReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	if _, err = s.client.Delete(ctx, converter.DeleteReq{req}.ConvertToProto()); err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	return nil, nil
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
