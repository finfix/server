package endpoint

import (
	"context"
	"net/http"
	"strconv"

	"server/app/pkg/errors"
	"server/app/pkg/server/middleware"
	"server/app/services"
	"server/app/services/transaction/model"
)

// @Summary Загрузка файла пользователя
// @Tags transaction
// @Security AuthJWT
// @Accept multipart/form-data
// @Param file formData file true "file to upload"
// @Success 200 {object} model.CreateFileRes
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /transaction/{transaction_id}/file [post]
func (s *endpoint) createFile(ctx context.Context, r *http.Request) (_ any, err error) {

	var req model.CreateFileReq

	transactionID, err := strconv.Atoi(r.PathValue("transaction_id"))
	if err != nil {
		return nil, errors.BadRequest.Wrap(err)
	}
	req.TransactionID = uint32(transactionID)

	// Читаем файл из запроса
	req.File, req.FileHeader, err = r.FormFile("file")
	if err != nil {
		return nil, errors.BadRequest.Wrap(err)
	}

	necessaryInformation, err := services.ExtractNecessaryFromCtx(ctx)
	if err != nil {
		return nil, errors.BadRequest.Wrap(err)
	}

	err = middleware.SetNecessary(necessaryInformation, &req)
	if err != nil {
		return nil, errors.BadRequest.Wrap(err)
	}

	id, err := s.service.CreateFile(ctx, req)
	if err != nil {
		return nil, err
	}

	return model.CreateFileRes{ID: id}, nil
}
