package osscontribute

import (
	"sync"
	"time"
)

type Cache struct {
	data      []byte
	timestamp time.Time
	mu        sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = data
	c.timestamp = time.Now().UTC()
}

func (c *Cache) Get() ([]byte, time.Time) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data, c.timestamp
}
