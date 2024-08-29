package cache

import "fmt"

// Custom types
type Key string
type Value interface{}

type Cache struct {
	items map[Key]Value
}

// Cunstructor
func New() *Cache {
	return &Cache{
		items: make(map[Key]Value),
	}
}

// Get a value by a key
func (c *Cache) Get(key Key) Value {
	value, exist := c.items[key]

	if !exist {
		fmt.Println("Key " + string(key) + " not found.")
		return nil
	}

	return value
}

// Adds or updates a key-value pair in cache
func (c *Cache) Set(key Key, value Value) {
	c.items[key] = value
}

// Deletes a key-value pair from the cache
func (c *Cache) Delete(key Key) {
	delete(c.items, key)
}
