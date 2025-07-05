package store

import "sync"

type Store interface {
	get(key string) (string, error)
	listKeys() ([]string, error)
	set(key string, value string) (bool, error)
	remove(key string)
}

type cache struct {
	data map[string]string
	mu   sync.Mutex
}

var c = cache{
	data: make(map[string]string),
}

func GetCache() *cache {
	return &c
}

func (c *cache) Get(key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.data[key]

	if !ok {
		return "No value found", nil
	}

	return v, nil
}

func (c *cache) ListKeys() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	keys := make([]string, 0, len(c.data))

	for k := range c.data {
		keys = append(keys, k)
	}

	return keys
}

func (c *cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

func (c *cache) Set(key string, value string) (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value

	return true, nil
}
