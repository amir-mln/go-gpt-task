package usecases

import "go-gpt-task/models"

type Usecases interface {
	ParsePrompt(prompt string)
}

type DbRepository interface {
	Insert(laptop models.Laptop)
	FindByID(id string) *models.Laptop
}

type CacheRepository interface {
	Contains(key string) interface{}
	Add(key string, data interface{})
}

type AIPromptParser interface {
	Parse(prompt string) (models.Laptop, error)
}
