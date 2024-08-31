package cache

import (
	"sync"
	"time"
)

type сacheItem[V any] struct {
	Value      V
	Expiration int64
}

type Cache[K comparable, V any] struct {
	mu         sync.RWMutex
	items      map[K]сacheItem[V]
	defaultTTL time.Duration
}

func NewCache[K comparable, V any](defaultTTL time.Duration) *Cache[K, V] {
	return &Cache[K, V]{
		items:      make(map[K]сacheItem[V]),
		defaultTTL: defaultTTL,
	}
}

func (c *Cache[K, V]) Set(key K, value V, ttl ...time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var expiration int64
	if len(ttl) > 0 {
		expiration = time.Now().Add(ttl[0]).UnixNano()
	} else {
		expiration = time.Now().Add(c.defaultTTL).UnixNano()
	}

	c.items[key] = сacheItem[V]{
		Value:      value,
		Expiration: expiration,
	}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found || time.Now().UnixNano() > item.Expiration {
		var zero V
		return zero, false
	}

	return item.Value, true
}
