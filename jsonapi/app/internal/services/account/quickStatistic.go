package account

import (
	"context"
	"net/http"

	"jsonapi/app/internal/services/account/converter"
	"jsonapi/app/internal/services/account/model"
	"pkg/errors"
	"pkg/validation"
)

// @Summary Получение краткой статистики по счетам
// @Tags account
// @Security AuthJWT
// @Param Query query model.QuickStatistic true "model.QuickStatisticReq"
// @Produce jsonapi
// @Success 200 {object} []model.QuickStatistic "[]model.QuickStatistic"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account/quickStatistic [get]
func (s *service) quickStatistic(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeQuickStatisticReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	statistic, err := s.client.QuickStatistic(ctx, converter.QuickStatisticReq{req}.ConvertToProto())
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}
	return converter.PbQuickStatisticRes{statistic}.ConvertToStruct().QuickStatistic, nil
}

func decodeQuickStatisticReq(ctx context.Context, _ *http.Request) (req model.QuickStatisticReq, err error) {

	// Заполняем поля из контекста
	req.UserID, _ = ctx.Value("UserID").(uint32)
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
