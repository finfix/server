package accountPermissions

import (
	"sync"
	"time"

	"server/pkg/sql"
)

type AccountPermissionsService struct {
	db          sql.SQL
	permissions permissions
}

var generalPermissions = AccountPermissions{
	UpdateBudget:          true,
	UpdateRemainder:       true,
	UpdateCurrency:        true,
	UpdateParentAccountID: true,

	CreateTransaction: true,
}

func NewAccountPermissionsService(db sql.SQL) (*AccountPermissionsService, error) {

	service := &AccountPermissionsService{
		db: db,
		permissions: permissions{
			typeToPermissions:     nil,
			isParentToPermissions: nil,
			mu:                    sync.RWMutex{},
		},
	}

	err := service.refreshAccountPermissions(true)
	if err != nil {
		return nil, err
	}
	go func() {
		time.Sleep(time.Minute)
		_ = service.refreshAccountPermissions(false)
	}()

	return service, nil
}
