package in_memory_cache

type Cache struct {
	key   string
	value interface{}
}

func NewCache() *Cache {
	return &Cache{}
}
func (cache *Cache) Set(key string, value interface{}) {
	cache.key = key
	cache.value = value
}
func (cache *Cache) Get(key string) interface{} {
	var value interface{}
	if key == cache.key {
		value = cache.value
	}
	return value
}
func (cache *Cache) Delete(key string) {
	if key == cache.key {
		cache.value = nil
	}
}
