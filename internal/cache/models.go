package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mu   sync.Mutex
	data map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
