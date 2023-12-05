package transaction

import (
	"context"
	"encoding/json"
	"net/http"

	"jsonapi/app/internal/services/transaction/converter"
	"jsonapi/app/internal/services/transaction/model"
	"pkg/errors"
	"pkg/validation"
)

// @Summary Создание транзакции
// @Description Создание транзакции и изменение баланса счетов, между которыми она произошла
// @Tags transaction
// @Security AuthJWT
// @Accept jsonapi
// @Param Body body model.CreateReq true "model.CreateReq"
// @Produce jsonapi
// @Success 200 {object} model.CreateRes
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /transaction [post]
func (s *service) createTransaction(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем тело запроса в структуру
	req, err := decodeCreateTransactionReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	proto, err := s.client.Create(ctx, converter.CreateReq{req}.ConvertToProto())
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Конвертируем ответ во внутреннюю структуру
	return converter.PbCreateRes{proto}.ConvertToStruct(), nil
}

func decodeCreateTransactionReq(ctx context.Context, r *http.Request) (req model.CreateReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Валидируем поля
	if err = req.Type.Validate(); err != nil {
		return req, err
	}
	if req.AmountFrom <= 0 || req.AmountTo <= 0 {
		return req, errors.BadRequest.New("amountFrom and amountTo must be greater than 0")
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
