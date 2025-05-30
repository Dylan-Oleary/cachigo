package store

type Store interface {
	get(key string) (string, error)
	set(key string, value string) (bool, error)
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

func (c *cache) Set(key string, value string) (bool, error) {
	c.data[key] = value

	return true, nil
}
