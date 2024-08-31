package accountPermissions

import (
	"sync"

	"server/internal/services/account/model/accountType"
)

type permissions struct {
	typeToPermissions     map[accountType.Type]AccountPermissions
	isParentToPermissions map[bool]AccountPermissions
	mu                    sync.RWMutex
}

func (p *permissions) get() (
	map[accountType.Type]AccountPermissions,
	map[bool]AccountPermissions,
) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.typeToPermissions, p.isParentToPermissions
}

func (p *permissions) set(
	typeToPermissions map[accountType.Type]AccountPermissions,
	isParentToPermissions map[bool]AccountPermissions,
) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.typeToPermissions = typeToPermissions
	p.isParentToPermissions = isParentToPermissions
}
