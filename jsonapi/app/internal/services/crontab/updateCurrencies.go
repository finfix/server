package crontab

import (
	"context"
	"net/http"

	"jsonapi/app/internal/services/crontab/converter"
	"pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) updateCurrencies(ctx context.Context, _ *http.Request) (any, error) {

	// Вызываем метод сервиса
	proto, err := s.client.UpdateCurrencies(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Конвертируем ответ во внутреннюю структуру
	return converter.PbUpdateCurrenciesRes{proto}.ConvertToStruct(), nil
}
