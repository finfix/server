package transaction

import (
	"context"
	"net/http"

	"jsonapi/app/internal/services/transaction/converter"

	"github.com/gorilla/schema"

	"jsonapi/app/internal/services/transaction/model"
	"pkg/errors"
	"pkg/validation"
)

// @Summary Получение всех транзакций
// @Description Получение всех транзакций по фильтрам
// @Tags transaction
// @Security AuthJWT
// @Param Query query model.GetReq true "model.CreateReq"
// @Produce json
// @Success 200 {object} []model.Transaction
// @Failure 400,404,500 {object} errors.CustomError
// @Router /transaction [get]
func (s *service) getTransactions(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем параметры запроса в структуру
	req, err := decodeGetTransactionsReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	proto, err := s.client.Get(ctx, converter.GetReq{req}.ConvertToProto())
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Конвертируем ответ во внутреннюю структуру
	return converter.PbGetRes{proto}.ConvertToStruct().Transactions, nil
}

func decodeGetTransactionsReq(ctx context.Context, r *http.Request) (req model.GetReq, err error) {

	// Декодируем параметры запроса в структуру
	if err = schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Валидируем поля
	if err = req.Type.Validate(); err != nil {
		return req, err
	}
	if req.DateFrom != nil && req.DateTo != nil {
		if req.DateFrom.After(req.DateTo.Time) || req.DateFrom.Equal(req.DateTo.Time) {
			return req, errors.BadRequest.New("date_from must be less than date_to")
		}
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
