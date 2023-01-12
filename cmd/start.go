package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/byyjoww/league-mentor/config"
	"github.com/byyjoww/league-mentor/logging"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts a new server",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := loadDevEnvs(); err != nil {
				panic(fmt.Errorf("failed to load dev environment variables: %w", err))
			}

			var (
				cfg          = config.Build()
				logrusLogger = logging.NewLogrusLogger(cfg.Logging)
			)

			logrusLogger.Info("Initializing app")

			switch startFlags.serverType {
			case serverTypeApi:
				startApi(logrusLogger, cfg)
			case serverTypeTerminal:
				startTerminal(logrusLogger, cfg)
			}
		},
	}

	startFlags = struct {
		serverType string
	}{}

	serverTypeApi      string = "api"
	serverTypeTerminal string = "terminal"
)

func init() {
	RootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&startFlags.serverType, "type", serverTypeApi, "The server type to initialize")
}

func loadDevEnvs() error {
	bytes, err := os.ReadFile("./config/dev-env.json")
	if err != nil {
		return err
	}

	var obj struct {
		ApiKey string
	}

	if err := json.Unmarshal(bytes, &obj); err != nil {
		return err
	}

	return os.Setenv("APP_CHATGPT_APIKEY", obj.ApiKey)
}
