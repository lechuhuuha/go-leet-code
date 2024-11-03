package cachemanagement

import (
	"sync"
	"time"
)

type CacheObject struct {
	Value string
	TTL   int64
}

func (c *CacheObject) Expired() bool {
	if c.TTL == 0 {
		return false
	}

	return time.Now().UnixNano() > c.TTL
}

type Cache struct {
	objects map[string]CacheObject
	mutex   *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		objects: map[string]CacheObject{},
		mutex:   &sync.RWMutex{},
	}
}

func (c *Cache) Get(key string) string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	obj := c.objects[key]
	if obj.Expired() {
		delete(c.objects, key)
	}

	return obj.Value
}

func (c *Cache) Set(key, value string, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.objects[key] = CacheObject{
		Value: value,
		TTL:   time.Now().Add(ttl).UnixNano(),
	}
}
