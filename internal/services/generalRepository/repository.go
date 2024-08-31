package generalRepository

import (
	"context"
	"sync"
	"time"

	"server/pkg/log"
	"server/pkg/sql"
)

type Repository struct {
	db       sql.SQL
	accesses accessesMap
}

func New(db sql.SQL) (_ *Repository, err error) {

	repository := &Repository{
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
