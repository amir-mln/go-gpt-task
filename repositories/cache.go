package repositories

import "go-gpt-task/usecases"

var (
	_ usecases.CacheRepository = Cache{}
)

type Cache map[string]usecases.CachedLaptopPrompt

func NewCache() Cache {
	return make(Cache, 0)
}

func (cache Cache) FindByKey(key string) (usecases.CachedLaptopPrompt, bool) {
	res, ok := cache[key]
	return res, ok
}

func (cache Cache) Insert(key string, data usecases.CachedLaptopPrompt) {
	cache[key] = data
}
