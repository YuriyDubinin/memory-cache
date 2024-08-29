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
	var value, isExist = c.items[key]

	if !isExist {
		fmt.Println("Key not found")
		return nil
	}

	return value
}
