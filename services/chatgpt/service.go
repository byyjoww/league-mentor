package chatgpt

import (
	"context"
	"strings"

	"github.com/byyjoww/league-mentor/config"
	gogpt "github.com/sashabaranov/go-gpt3"
)

const (
	Model       = gogpt.GPT3TextDavinci003
	MaxTokens   = 2048
	Temperature = 0.1
	TopP        = 1
	N           = 1
	LogProbs    = 0
)

type Client interface {
	CreateCompletion(ctx context.Context, prompt string) (Completion, error)
}

type GoGptClient struct {
	client *gogpt.Client
}

func NewGoGptClient(cfg config.ChatGPT) *GoGptClient {
	return &GoGptClient{
		client: gogpt.NewClient(cfg.ApiKey),
	}
}

func (c *GoGptClient) CreateCompletion(ctx context.Context, prompt string) (Completion, error) {
	req := gogpt.CompletionRequest{
		Model:       Model,
		MaxTokens:   MaxTokens,
		Prompt:      prompt,
		Temperature: Temperature,
		Stream:      false,
		TopP:        TopP,
		N:           N,
		LogProbs:    LogProbs,
		Stop:        []string{"/n"},
	}

	resp, err := c.client.CreateCompletion(ctx, req)
	if err != nil {
		return Completion{}, err
	}

	return Completion{
		Value: formatCompletion(resp.Choices[0].Text),
	}, nil
}

func formatCompletion(text string) string {
	return strings.TrimLeft(text, "\n")
}
