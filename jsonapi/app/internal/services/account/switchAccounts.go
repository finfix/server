package account

import (
	"context"
	pb "core/app/proto/pbAccount"
	"encoding/json"
	"net/http"
	"pkg/converter"

	"jsonapi/app/internal/services/account/model"
	"pkg/errors"
	"pkg/validation"
)

// @Summary Изменение порядковых мест двух счетов
// @Description Поменять два счета местами
// @Tags account
// @Security AuthJWT
// @Accept json
// @Param Body body model.SwitchReq true "model.SwitchReq"
// @Produce json
// @Success 200 "Если изменение порядка счетов прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account/switch [patch]
func (s *service) switchAccounts(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeSwitchAccountsReq(ctx, r)
	if err != nil {
		return nil, err
	}

	in, err := converter.Convert(pb.SwitchReq{}, req)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	if _, err = s.client.Switch(ctx, &in); err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	return nil, nil
}

func decodeSwitchAccountsReq(ctx context.Context, r *http.Request) (req model.SwitchReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
