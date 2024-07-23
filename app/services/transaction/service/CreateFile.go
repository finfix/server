package service

import (
	"context"
	"io"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"

	"server/app/pkg/errors"
	"server/app/services/generalRepository/checker"
	transactionModel "server/app/services/transaction/model"
	transactionRepoModel "server/app/services/transaction/repository/model"
)

// CreateFile сохраняет файл
func (s *Service) CreateFile(ctx context.Context, req transactionModel.CreateFileReq) (id uint32, err error) {

	defer req.File.Close()

	// Проверяем, имеет ли пользователь доступ к транзакции
	if err = s.generalRepository.CheckUserAccessToObjects(ctx, checker.Transactions, req.Necessary.UserID, []uint32{req.TransactionID}); err != nil {
		return id, err
	}

	// Разбиваем имя файла на название и расширение
	filenameArr := strings.Split(req.FileHeader.Filename, ".")
	if len(filenameArr) != 2 { //nolint:gomnd
		return id, errors.BadRequest.New("Файл не имеет расширения")
	}

	// Генерируем новое уникальное имя файла
	uniqueFileName := uuid.New().String() + "." + filenameArr[1]
	fileUrl := "files/" + uniqueFileName

	// Создаем файл на машине
	file, err := os.Create(fileUrl)
	if err != nil {
		return id, errors.InternalServer.Wrap(err)
	}
	defer file.Close()

	// Сохраняем файл от юзера на машину
	if _, err = io.Copy(file, req.File); err != nil {
		return id, errors.InternalServer.Wrap(err)
	}

	// Сохраняем данные о файле в базу данных
	id, err = s.transactionRepository.CreateFile(ctx, transactionRepoModel.CreateFileReq{
		TransactionID:    req.TransactionID,
		OriginalFileName: req.FileHeader.Filename,
		UniqueFileName:   uniqueFileName,
		DatetimeCreate:   time.Now(),
		CreatedByUserID:  req.Necessary.UserID,
	})
	if err != nil {
		return id, err
	}

	return id, nil
}
