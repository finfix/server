package endpoint

import (
	"context"
	"net/http"

	"pkg/http/decoder"

	"server/internal/services/accountGroup/model"
)

// @Summary Удаление группы счетов
// @Tags accountGroup
// @Security AuthJWT
// @Param Query query model.DeleteAccountGroupReq true "model.DeleteAccountGroupReq"
// @Produce json
// @Success 200 "Если удаление группы счетов прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /accountGroup [delete]
func (s *endpoint) deleteAccountGroup(ctx context.Context, r *http.Request) (any, error) {

	var req model.DeleteAccountGroupReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.DeleteAccountGroup(ctx, req)
}
