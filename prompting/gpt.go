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
	systemMessage = `
	You are an expert in identifying and extracting hardware information from laptops.
	Given input data, your task is to extract information based on the following schema:
		- Message: A human-readable message when the input is non-laptop related, invalid or doesn't contain enough information about laptops.
		- Failed: A boolean indicating whether data extraction succeeded.
		- Laptop: Contains extracted laptop data if successful, or null if no data could be extracted.
	You need to consider the following notes when it comes to extracting laptop information:
		- if a field's data is not specified, use empty strings ("")
		- the model and the brand of the data should not be mistaken with each other. if one is provided
		  and the other is not, search the web to find out
		- we need information about processor's manufacturer and its model. you can search the web if the
		  provided user's message doesn't contain the necessary information. but it's important to match
		  the extracted processor information with search results.
		- if the type of the storage is not specified, it's safe to assume it is HDD
		- do not round the numbers.
		- it's important to not add inaccurate information from the web search result. if searching
		  does not provide you with information that relates to searched field, ignore it.
	`
	assistantMessage = "The provided input does not relate to laptops"
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
					Role:    openai.ChatMessageRoleAssistant,
					Content: assistantMessage,
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

func NewGPTPromptParser(key string) *GPTPromptParser {
	return &GPTPromptParser{
		client: openai.NewClient(key),
	}
}
