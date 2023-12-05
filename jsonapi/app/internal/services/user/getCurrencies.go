package user

import (
	"context"
	"net/http"

	"jsonapi/app/internal/services/user/converter"
	_ "jsonapi/app/internal/services/user/model"
	"pkg/errors"

	"google.golang.org/protobuf/types/known/emptypb"
)

// @Summary Получение списка валют
// @Tags user
// @Security AuthJWT
// @Produce jsonapi
// @Success 200 {object} []model.Currency
// @Failure 401,500 {object} errors.CustomError
// @Router /user/getCurrencies [get]
func (s *service) getCurrencies(ctx context.Context, _ *http.Request) (any, error) {

	// Вызываем метод сервиса
	res, err := s.client.GetCurrencies(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Конвертируем ответ во внутреннюю структуру
	return converter.PbGetCurrenciesRes{res}.ConvertToStruct().Currencies, nil
}
