package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/byyjoww/league-mentor/config"
	"github.com/byyjoww/league-mentor/services/chatgpt"
	"github.com/sirupsen/logrus"
)

func startTerminal(logrusLogger logrus.FieldLogger, cfg config.Config) {
	gptClient := chatgpt.NewGoGptClient(cfg.ChatGPT)
	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()
	var prompt string

	printToStdOut("Ask your question: ")
	for scanner.Scan() {
		prompt = scanner.Text()

		completion, err := gptClient.CreateCompletion(ctx, prompt)
		if err != nil {
			logrusLogger.WithError(err).Error("failed to create completion")
			return
		}

		printToStdOut(completion.Value)
		printToStdOut("Ask your question: ")
	}

	if err := scanner.Err(); err != nil {
		logrusLogger.WithError(err).Fatal("error scanning stdout")
	}
}

func printToStdOut(msg string) {
	fmt.Println(msg)
}
