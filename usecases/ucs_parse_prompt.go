package usecases

import (
	"errors"
	"fmt"
	"go-gpt-task/models"

	"github.com/google/uuid"
)

type PromptState int

const (
	Successful PromptState = iota + 1
	Failed
)

type PromptCacheResult struct {
	Value string
	State PromptState
}

func (uc *usecases) ParsePrompt(prompt string) (*models.Laptop, error) {
	data := uc.cacheRepo.Contains(prompt)
	if result, ok := data.(PromptCacheResult); ok {
		if result.State == Failed {
			return nil, errors.New(result.Value)
		}

		laptop := uc.dbRepo.FindByID(result.Value)
		if laptop == nil {
			return nil, fmt.Errorf(
				"for prompt %q, found a non existing cached laptop with id %q",
				prompt,
				result.Value,
			)

		}

		return laptop, nil
	} else if data != nil {
		return nil, fmt.Errorf("for prompt %q, found a cached result with invalid type %T", prompt, data)
	}

	laptop, err := uc.aiParser.Parse(prompt)
	if err != nil {
		result := PromptCacheResult{
			State: Failed,
			Value: err.Error(),
		}
		uc.cacheRepo.Add(prompt, result)

		return nil, err
	}

	laptop.ID = uuid.NewString()
	if err := laptop.Validate(); err != nil {
		newErr := fmt.Errorf("while parsing prompt:\n%q,\nfound following error(s): %w", prompt, err)
		result := PromptCacheResult{
			State: Failed,
			Value: newErr.Error(),
		}
		uc.cacheRepo.Add(prompt, result)

		return nil, err
	}

	uc.cacheRepo.Add(prompt, PromptCacheResult{State: Successful, Value: laptop.ID})
	uc.dbRepo.Insert(laptop)

	return &laptop, nil
}
