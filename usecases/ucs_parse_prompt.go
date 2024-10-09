package usecases

import (
	"context"
	"errors"
	"fmt"
	"go-gpt-task/models"
	"strings"

	"github.com/google/uuid"
)

type CachedLaptopPrompt struct {
	Prompt string
	Value  string
	Failed bool
}

type LaptopPromptSchema struct {
	Message string        `json:"message"`
	Failed  bool          `json:"failed"`
	Laptop  models.Laptop `json:"laptop"`
}

func (uc *usecases) ParsePrompt(ctx context.Context, prompt string) (models.Laptop, error) {
	if strings.Trim(prompt, " ") == "" {
		return models.Laptop{}, errors.New("invalid empty string as prompt")
	}

	if cached, ok := uc.cacheRepo.FindByKey(prompt); ok {
		if cached.Failed {
			return models.Laptop{}, errors.New(cached.Value)
		}

		if laptop, ok := uc.dbRepo.FindByID(cached.Value); !ok {
			return models.Laptop{}, fmt.Errorf(
				"found a non existing cached laptop with id %q",
				cached.Value,
			)
		} else {
			return laptop, nil
		}

	}

	if resp, err := uc.aiParser.Parse(ctx, prompt); err == nil {
		if resp.Failed {
			uc.cacheRepo.Insert(
				prompt,
				CachedLaptopPrompt{Failed: true, Value: resp.Message, Prompt: prompt},
			)
			return models.Laptop{}, errors.New(resp.Message)
		}

		laptop := resp.Laptop
		laptop.ID = uuid.NewString()
		if err := laptop.Validate(); err != nil {
			uc.cacheRepo.Insert(
				prompt,
				CachedLaptopPrompt{Failed: true, Prompt: prompt, Value: err.Error()},
			)

			return models.Laptop{}, err
		}

		uc.cacheRepo.Insert(
			prompt,
			CachedLaptopPrompt{Failed: false, Value: laptop.ID, Prompt: prompt},
		)
		uc.dbRepo.Insert(laptop)

		return laptop, nil
	} else {
		return models.Laptop{}, err
	}
}
