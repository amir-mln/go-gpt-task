package repositories

import "go-gpt-task/usecases"

var (
	_ usecases.CacheRepository = Cache{}
)

type Cache map[string]usecases.CachedLaptopPrompt

func (cache Cache) FindByKey(key string) (usecases.CachedLaptopPrompt, bool) {
	data := cache[key]
	return data, true
}

func (cache Cache) Insert(key string, data usecases.CachedLaptopPrompt) {
	cache[key] = data
}
