package memcache

import (
	"sync"
	"time"
)

var cache = NewCache()

type item struct {
	value  any
	expire time.Time
}

type memCache struct {
	mu    sync.Mutex
	store map[string]item
}

func (m *memCache) Set(key string, value any, ttl time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var t time.Time
	if ttl != 0 {
		t = time.Now().Add(ttl)
	}

	m.store[key] = item{value: value, expire: t}
}

func (m *memCache) Get(key string) (any, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	item, ok := m.store[key]
	if !ok {
		return nil, false
	}

	if item.expire.IsZero() || item.expire.After(time.Now()) {
		return item.value, ok
	}

	delete(m.store, key)
	return nil, false
}

func (m *memCache) Remove(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.store, key)
}

func NewCache() *memCache {
	return &memCache{
		mu:    sync.Mutex{},
		store: map[string]item{},
	}
}
