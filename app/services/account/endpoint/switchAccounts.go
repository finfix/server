package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/pkg/errors"
	"server/app/pkg/validation"
	"server/app/services"
	"server/app/services/account/model"
)

// @Summary Изменение порядковых мест двух счетов
// @Description Поменять два счета местами
// @Tags account
// @Security AuthJWT
// @Accept json
// @Param Body body model.SwitchAccountBetweenThemselvesReq true "model.SwitchAccountBetweenThemselvesReq"
// @Produce json
// @Success 200 "Если изменение порядка счетов прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account/switch [patch]
func (s *endpoint) switchAccounts(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeSwitchAccountsReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.SwitchAccountBetweenThemselves(ctx, req)
}

func decodeSwitchAccountsReq(ctx context.Context, r *http.Request) (req model.SwitchAccountBetweenThemselvesReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.Necessary, err = services.ExtractNecessaryFromCtx(ctx)
	if err != nil {
		return req, err
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
