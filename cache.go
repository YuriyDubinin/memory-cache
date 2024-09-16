package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var errKeyNotFound = errors.New("key not found")
var errLifetimeExpired = errors.New("life time has expire, the key was deleted")

type Key string
type Value interface{}

type CacheItem struct {
	value          Value
	expirationTime time.Time
}
type Cache struct {
	storage map[Key]CacheItem
	mutex   sync.RWMutex
}

func New() *Cache {
	return &Cache{
		storage: make(map[Key]CacheItem),
	}
}

func (c *Cache) Get(key Key) (Value, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	currentTime := time.Now()

	value, exist := c.storage[key]
	if !exist {
		return nil, fmt.Errorf("%w: %v", errKeyNotFound, key)
	}
	if currentTime.After(c.storage[key].expirationTime) {
		delete(c.storage, key)
		return nil, fmt.Errorf("%w", errLifetimeExpired)
	}

	return value, nil
}

func (c *Cache) Set(key Key, value Value, lifetime time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expirationTime := time.Now().Add(time.Second * lifetime)
	c.storage[key] = CacheItem{value: value, expirationTime: expirationTime}
}

func (c *Cache) Delete(key Key) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, exist := c.storage[key]
	if !exist {
		return fmt.Errorf("%w: %v", errKeyNotFound, key)
	}

	delete(c.storage, key)

	return nil
}
