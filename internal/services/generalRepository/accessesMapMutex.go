package generalRepository

import "sync"

type accessesMap struct {
	accesses map[uint32]map[uint32]struct{}
	mu       sync.RWMutex
}

func (am *accessesMap) Get() map[uint32]map[uint32]struct{} {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return am.accesses
}

func (am *accessesMap) Set(accesses map[uint32]map[uint32]struct{}) {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.accesses = accesses
}
