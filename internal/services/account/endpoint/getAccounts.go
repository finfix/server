package endpoint

import (
	"context"
	"net/http"

	"server/pkg/http/decoder"
	"server/internal/services/account/model"
)

// @Summary Получение счетов по фильтрам
// @Tags account
// @Security AuthJWT
// @Param Query query model.GetAccountsReq false "model.GetAccountsReq"
// @Produce json
// @Success 200 {object} []model.Account
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /account [get]
func (s *endpoint) get(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetAccountsReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetAccounts(ctx, req)
}
