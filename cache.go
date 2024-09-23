package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var errKeyNotFound = errors.New("key not found")

type Key string
type Value interface{}

type CacheItem struct {
	value          Value
	expirationTime time.Time
}

type Cache struct {
	storage        map[Key]CacheItem
	mutex          sync.RWMutex
	stopClearingCh chan bool
}

func New() *Cache {
	cache := &Cache{
		storage:        make(map[Key]CacheItem),
		stopClearingCh: make(chan bool),
	}

	go cache.clearOutdatedData()

	return cache
}

func (c *Cache) Get(key Key) (Value, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, exist := c.storage[key]
	if !exist {
		return nil, fmt.Errorf("%w: %v", errKeyNotFound, key)
	}

	return value.value, nil
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

func (c *Cache) clearOutdatedData() {
	ticker := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-ticker.C:
			for key, item := range c.storage {
				if time.Now().After(item.expirationTime) {
					c.Delete(key)
					c.stopDataClearing()
				}
			}
		case <-c.stopClearingCh:
			ticker.Stop()
			return
		}

	}

}

func (c *Cache) stopDataClearing() {
	c.stopClearingCh <- true
}
