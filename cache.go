package cache

type Cache struct {
	items map[string]interface{}
}

func New() Cache {
	return Cache{
		items: make(map[string]interface{}),
	}
}
