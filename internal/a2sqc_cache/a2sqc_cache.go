package a2sqc_cache

import (
	"encoding/hex"
	"strings"
	"sync"
	"time"
)

type Cache struct {
	data  map[string][]byte
	mutex sync.RWMutex
	ttl   time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		data: make(map[string][]byte),
		ttl:  ttl,
	}
}

func (c *Cache) Set(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value

	time.AfterFunc(c.ttl, func() {
		c.mutex.Lock()
		delete(c.data, key)
		c.mutex.Unlock()
	})
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	data, exists := c.data[key]
	return data, exists
}

var cacheableQueries = []string{
	"FFFFFFFF54", // A2S_INFO
	"FFFFFFFF55", // A2S_PING
	"FFFFFFFF41", // A2S_PLAYER
}

func IsCacheable(query []byte) bool {
	hexQuery := strings.ToUpper(hex.EncodeToString(query))
	for _, prefix := range cacheableQueries {
		if strings.HasPrefix(hexQuery, prefix) {
			return true
		}
	}
	return false
}
