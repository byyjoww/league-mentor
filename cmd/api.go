package cmd

import (
	"context"
	"fmt"

	"github.com/byyjoww/league-mentor/config"
	"github.com/byyjoww/league-mentor/services/chatgpt"
	"github.com/sirupsen/logrus"
)

func startApi(logrusLogger logrus.FieldLogger, cfg config.Config) {
	ctx := context.Background()
	prompt := "How do I beat malphite in lane?"

	gptClient := chatgpt.NewGoGptClient(cfg.ChatGPT)
	completion, err := gptClient.CreateCompletion(ctx, prompt)
	if err != nil {
		logrusLogger.WithError(err).Error("failed to create completion")
		return
	}

	fmt.Println(completion.Value)
}
