package usecases

type usecases struct {
	dbRepo    DbRepository
	cacheRepo CacheRepository
	aiParser  AIPromptParser
}

func NewUsecases(db DbRepository, cache CacheRepository, parser AIPromptParser) *usecases {
	uc := &usecases{
		db,
		cache,
		parser,
	}

	return uc
}
