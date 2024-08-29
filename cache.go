package cache

import "fmt"

type Key string
type Value interface{}

type Cache struct {
	items map[Key]Value
}

func New() *Cache {
	return &Cache{
		items: make(map[Key]Value),
	}
}

func (c *Cache) Get(key Key) Value {
	value, exist := c.items[key]

	if !exist {
		fmt.Println("Key " + string(key) + " not found.")
		return nil
	}

	return value
}

func (c *Cache) Set(key Key, value Value) {
	c.items[key] = value
}

func (c *Cache) Delete(key Key) {
	delete(c.items, key)
}
