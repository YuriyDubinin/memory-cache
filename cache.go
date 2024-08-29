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

func (c *Cache) Set(key Key, value Value) {
	c.items[key] = value
	fmt.Println("The key " + string(key) + " successfully recorded.")
}
