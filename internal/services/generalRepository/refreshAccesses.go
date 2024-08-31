package generalRepository

import (
	"context"
	"time"

	"server/pkg/log"
)

func (repo *GeneralRepository) refreshAccesses(doOnce bool) error {
	for {
		log.Info(context.Background(), "Получаем доступы пользователей к объектам")
		_accesses, err := repo.getAccesses(context.Background())
		repo.accesses.Set(_accesses)
		if doOnce {
			return err
		}
		if err != nil {
			log.Error(context.Background(), err)
		}

		time.Sleep(time.Minute)
	}
}
