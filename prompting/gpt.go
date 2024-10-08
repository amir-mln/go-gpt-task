package prompting

import (
	"context"
	"go-gpt-task/usecases"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

var (
	_ usecases.AIPromptParser = &GPTPromptParser{}
)

const (
	systemMessage = `You are an expert in identifying and extracting hardware information from laptops.
	Given input data, your task is to extract information based on the following schema:
		- Message: A human-readable message when the input is non-laptop related, invalid or doesn't contain enough information about laptops.
		- Failed: A boolean indicating whether data extraction succeeded.
		- Laptop: Contains extracted laptop data if successful, or null if no data could be extracted.
	`
)

type GPTPromptParser struct {
	client *openai.Client
}

func (gpt *GPTPromptParser) Parse(ctx context.Context, prompt string) (usecases.LaptopPromptSchema, error) {
	var response usecases.LaptopPromptSchema

	schema, err := jsonschema.GenerateSchemaForType(response)
	if err != nil {
		return response, err
	}

	resp, err := gpt.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemMessage,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
				JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:   "laptop_parser",
					Schema: schema,
					Strict: true,
				},
			},
		},
	)
	if err != nil {
		return response, err
	}

	err = schema.Unmarshal(resp.Choices[0].Message.Content, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
