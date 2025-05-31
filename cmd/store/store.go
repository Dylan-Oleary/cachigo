package store

type Store interface {
	get(key string) (string, error)
	listKeys() ([]string, error)
	set(key string, value string) (bool, error)
	remove(key string)
}

type cache struct {
	data map[string]string
}

var c = cache{
	data: make(map[string]string),
}

func InitCache() *cache {
	return &c
}

func (c *cache) Get(key string) (string, error) {
	v, ok := c.data[key]

	if !ok {
		return "No value found", nil
	}

	return v, nil
}

func (c *cache) ListKeys() []string {
	keys := make([]string, 0, len(c.data))

	for k := range c.data {
		keys = append(keys, k)
	}

	return keys
}

func (c *cache) Remove(key string) {
	delete(c.data, key)
}

func (c *cache) Set(key string, value string) (bool, error) {
	c.data[key] = value

	return true, nil
}
