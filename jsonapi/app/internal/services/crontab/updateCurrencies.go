package crontab

import (
	"context"
	"jsonapi/app/internal/services/crontab/model"
	"net/http"
	"pkg/converter"

	"pkg/errors"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) updateCurrencies(ctx context.Context, _ *http.Request) (any, error) {

	// Вызываем метод сервиса
	out, err := s.client.UpdateCurrencies(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	res, err := converter.Convert(model.UpdateCurrenciesRes{}, out)

	// Конвертируем ответ во внутреннюю структуру
	return res, nil
}
