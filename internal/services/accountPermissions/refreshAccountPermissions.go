package accountPermissions

import (
	"context"
	"time"

	"server/pkg/log"
)

func (s *Service) refreshAccountPermissions(doOnce bool) error {
	for {
		log.Info(context.Background(), "Получаем пермишены на действия со счетами")
		var err error
		_typeToPermissions, _isParentToPermissions, err := s.getAccountPermissions(context.Background())
		s.permissions.set(_typeToPermissions, _isParentToPermissions)
		if doOnce {
			return err
		}
		if err != nil {
			log.Error(context.Background(), err)
		}

		time.Sleep(time.Minute)
	}
}
