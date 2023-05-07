package aop

var inMemoryCache = map[string]interface{}{}

func GetCache[T any](key string) (*T, bool) {
	value, ok := inMemoryCache[key]
	if !ok {
		return nil, false
	}

	result, ok := value.(*T)
	if !ok {
		return nil, false
	}

	return result, true
}

func SetCache(key string, value interface{}) error {
	inMemoryCache[key] = value
	return nil
}

func ClearCache(key string) error {
	delete(inMemoryCache, key)
	return nil
}
