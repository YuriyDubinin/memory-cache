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
	value, isExist := c.items[key]

	if !isExist {
		fmt.Println("Key " + string(key) + " not found.")
		return nil
	}

	fmt.Println("The value for the key " + string(key) + " was successfully retrieved.")
	return value
}
