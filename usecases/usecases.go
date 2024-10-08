package usecases

import (
	"context"
	"go-gpt-task/models"
)

var (
	_ Usecases = &usecases{}
)

type Usecases interface {
	ParsePrompt(ctx context.Context, prompt string) (models.Laptop, error)
}

type DbRepository interface {
	Insert(laptop models.Laptop)
	FindByID(id string) (models.Laptop, bool)
}

type CacheRepository interface {
	Insert(key string, data CachedLaptopPrompt)
	FindByKey(key string) (CachedLaptopPrompt, bool)
}

type AIPromptParser interface {
	Parse(ctx context.Context, prompt string) (LaptopPromptSchema, error)
}

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
