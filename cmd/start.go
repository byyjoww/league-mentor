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
			switch startFlags.serverType {
			case serverTypeDevApi:
				loadDevEnvs()
			case serverTypeDevClient:
				loadDevEnvs()
			}

			var (
				cfg          = config.Build()
				logrusLogger = logging.NewLogrusLogger(cfg.Logging)
			)

			logrusLogger.Info("Initializing app")

			switch startFlags.serverType {
			case serverTypeApi:
				startApi(logrusLogger, cfg)
			case serverTypeClient:
				startClient(logrusLogger, cfg)
			case serverTypeDevApi:
				startApi(logrusLogger, cfg)
			case serverTypeDevClient:
				startClient(logrusLogger, cfg)
			}
		},
	}

	startFlags = struct {
		serverType string
	}{}

	serverTypeApi    string = "api"
	serverTypeClient string = "client"

	serverTypeDevApi    string = "dev-api"
	serverTypeDevClient string = "dev-client"
)

func init() {
	RootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&startFlags.serverType, "type", serverTypeApi, "The server type to initialize")
}

func loadDevEnvs() {
	bytes, err := os.ReadFile("./config/dev-env.json")
	if err != nil {
		err = fmt.Errorf("failed to load dev environment variables: %w", err)
		panic(err)
	}

	var obj struct {
		ChatGptApiKey   string `json:"chatGptApiKey"`
		RiotGamesApiKey string `json:"riotGamesApiKey"`
	}

	if err := json.Unmarshal(bytes, &obj); err != nil {
		err = fmt.Errorf("failed to load dev environment variables: %w", err)
		panic(err)
	}

	if err := os.Setenv("APP_CHATGPT_APIKEY", obj.ChatGptApiKey); err != nil {
		err = fmt.Errorf("failed to load dev environment variables: %w", err)
		panic(err)
	}

	if err := os.Setenv("APP_RIOTGAMES_APIKEY", obj.RiotGamesApiKey); err != nil {
		err = fmt.Errorf("failed to load dev environment variables: %w", err)
		panic(err)
	}
}
