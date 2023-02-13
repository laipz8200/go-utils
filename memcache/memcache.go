package memcache

import (
	"sync"
	"time"
)

var cache = &memcache{
	store: map[string]item{},
}

type item struct {
	value any
	ttl   int64
}

type memcache struct {
	mu    sync.Mutex
	store map[string]item
}

func (m *memcache) setWithTTL(key string, value any, ttl int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if ttl != 0 {
		ttl = time.Now().Unix() + ttl
	}

	m.store[key] = item{value: value, ttl: ttl}
}

func (m *memcache) get(key string) (any, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	item, ok := m.store[key]
	if !ok {
		return nil, false
	}

	if item.ttl == 0 || item.ttl > time.Now().Unix() {
		return item.value, ok
	}

	delete(m.store, key)
	return nil, false
}

func (m *memcache) remove(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.store, key)
}
