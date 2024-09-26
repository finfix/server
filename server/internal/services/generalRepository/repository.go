package generalRepository

import (
	"sync"
	"time"

	"pkg/sql"
)

type GeneralRepository struct {
	db       sql.SQL
	accesses accessesMap
}

func NewGeneralRepository(db sql.SQL) (_ *GeneralRepository, err error) {

	repository := &GeneralRepository{
		db: db,
		accesses: accessesMap{
			accesses: nil,
			mu:       sync.RWMutex{},
		},
	}

	err = repository.refreshAccesses(true)
	if err != nil {
		return nil, err
	}
	go func() {
		time.Sleep(time.Minute)
		_ = repository.refreshAccesses(false)
	}()

	return repository, nil
}
