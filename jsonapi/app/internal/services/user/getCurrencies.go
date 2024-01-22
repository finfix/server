package user

import (
	"context"
	"jsonapi/app/internal/services/user/model"
	"net/http"
	"pkg/converter"

	_ "jsonapi/app/internal/services/user/model"
	"pkg/errors"

	"google.golang.org/protobuf/types/known/emptypb"
)

// @Summary Получение списка валют
// @Tags user
// @Security AuthJWT
// @Produce json
// @Success 200 {object} []model.Currency
// @Failure 401,500 {object} errors.CustomError
// @Router /user/getCurrencies [get]
func (s *service) getCurrencies(ctx context.Context, _ *http.Request) (any, error) {

	// Вызываем метод сервиса
	out, err := s.client.GetCurrencies(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	res, err := converter.Convert(model.GetCurrenciesRes{}, out)
	if err != nil {
		return nil, err
	}

	// Конвертируем ответ во внутреннюю структуру
	return res.Currencies, nil
}
