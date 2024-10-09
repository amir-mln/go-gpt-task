package usecases_test

import (
	"context"
	"go-gpt-task/models"
	"go-gpt-task/usecases"
	"os"
	"testing"
)

type MockCache map[string]usecases.CachedLaptopPrompt

func (cache MockCache) FindByKey(key string) (usecases.CachedLaptopPrompt, bool) {
	res, ok := cache[key]
	return res, ok
}

func (cache MockCache) Insert(key string, data usecases.CachedLaptopPrompt) {
	cache[key] = data
}

type MockDatabase []models.Laptop

func (db *MockDatabase) Insert(laptop models.Laptop) {
	*db = append(*db, laptop)
}

func (db *MockDatabase) FindByID(id string) (models.Laptop, bool) {
	for _, lp := range *db {
		if lp.ID == id {
			return lp, true
		}
	}

	return models.Laptop{}, false
}

type ParserCounter struct {
	count uint
}

func (parser *ParserCounter) Parse(ctx context.Context, prompt string) (usecases.LaptopPromptSchema, error) {
	parser.count++
	return usecases.LaptopPromptSchema{}, nil
}

type NoopParser struct{}

func (parser *NoopParser) Parse(ctx context.Context, prompt string) (usecases.LaptopPromptSchema, error) {
	return usecases.LaptopPromptSchema{}, nil
}

var (
	db           usecases.DbRepository
	cache        usecases.CacheRepository
	parseCounter *ParserCounter
	noopParser   *NoopParser
)

func initServices() {
	db = &MockDatabase{}
	cache = MockCache{}
	parseCounter = &ParserCounter{}
	noopParser = &NoopParser{}
}

func TestMain(m *testing.M) {
	initServices()
	code := m.Run()
	os.Exit(code)
}

func TestRepetitiveInputs(t *testing.T) {
	ucs := usecases.NewUsecases(db, cache, parseCounter)

	cases := []string{
		"MacBook Pro with M1 chip, 8GB RAM, 256 GB SSD storage Battery removed",
		"MacBook Pro with M1 chip, 8GB RAM, 256 GB SSD storage Battery removed",
		"MacBook Pro with M1 chip, 8GB RAM, 256 GB SSD storage Battery removed",
	}

	for _, c := range cases {
		ucs.ParsePrompt(context.Background(), c)
	}

	if parseCounter.count != 1 {
		t.Errorf("for 3 identical inputs, AIPromptParser service was called more than once")
	}

	t.Cleanup(initServices)
}

func TestEmptyInput(t *testing.T) {
	ucs := usecases.NewUsecases(db, cache, noopParser)
	_, err := ucs.ParsePrompt(context.Background(), "   ")

	if err == nil {
		t.Errorf("was expecting to receive an error for empty string input")
	}

	t.Cleanup(initServices)
}
