package cache

import (
	"errors"
	"fmt"
	"sync"
)

// Custom types
type Key string
type Value interface{}

// Errors
var errKeyNotFound = errors.New("key not found")

type Cache struct {
	storage map[Key]Value
	mutex   sync.RWMutex
}

// Cunstructor
func New() *Cache {
	return &Cache{
		storage: make(map[Key]Value),
	}
}

// Get a value by a key
func (c *Cache) Get(key Key) (Value, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, exist := c.storage[key]

	if !exist {
		return nil, fmt.Errorf("%w: %s", errKeyNotFound, key)
	}

	return value, nil
}

// Adds or updates a key-value pair in cache
func (c *Cache) Set(key Key, value Value) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	c.storage[key] = value
}

// Deletes a key-value pair from the cache
func (c *Cache) Delete(key Key) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	_, exist := c.storage[key]

	if !exist {
		return fmt.Errorf("%w: %s", errKeyNotFound, key)
	}

	delete(c.storage, key)

	return nil
}
