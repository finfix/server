package service

import (
	"context"

	"server/app/services/accountGroup/model"
)

// CreateAccountGroup создает новую группу счетов
func (s *Service) CreateAccountGroup(ctx context.Context, accountGroup model.CreateAccountGroupReq) (res model.CreateAccountGroupRes, err error) {

	// Создаем SQL-транзакцию
	return res, s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Создаем счет
		if res.ID, res.SerialNumber, err = s.accountGroupRepository.CreateAccountGroup(ctx, accountGroup.ConvertToRepoReq()); err != nil {
			return err
		}

		if err = s.accountGroupRepository.LinkUserToAccountGroup(ctx, accountGroup.Necessary.UserID, res.ID); err != nil {
			return err
		}

		return nil
	})
}
