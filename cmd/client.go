package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/byyjoww/league-mentor/config"
	"github.com/byyjoww/league-mentor/servers/api/handler"
	"github.com/byyjoww/league-mentor/services/http/client"
	"github.com/byyjoww/league-mentor/services/http/client/auth"
	"github.com/byyjoww/league-mentor/services/http/decoder"
	"github.com/byyjoww/league-mentor/services/riotgames/lol"
	"github.com/sirupsen/logrus"
)

var (
	tickInterval        = 20 * time.Second
	summonerName string = "byyjoww"
)

func startClient(logrusLogger logrus.FieldLogger, cfg config.Config) {
	var (
		summoner lol.Summoner
		// overview string
	)

	restClient := client.New(http.DefaultClient, decoder.New()).
		WithAuth(auth.NewBasic("user", "pass")).
		WithBaseUrl("http://0.0.0.0:8080" + "/api/v1")

	statusCode, err := restClient.Get("/identity").
		WithContext(context.Background()).
		WithJsonBody(handler.IdentityRequest{SummonerName: summonerName}).
		DoAndUnmarshal(&summoner)

	if err != nil {
		logrusLogger.WithError(err).Fatal("failed to get identity")
		return
	}

	if statusCode >= 300 {
		err = fmt.Errorf("request failed with status code: %d", statusCode)
		logrusLogger.WithError(err).Fatal("failed to get identity")
		return
	}

	logger := logrusLogger.WithFields(logrus.Fields{
		"summoner": summoner,
	})

	ticker := time.NewTicker(tickInterval)
	done := make(chan bool)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			checkLaneOverview(logger, restClient, summoner)
		}
	}
}

func checkLaneOverview(logrusLogger logrus.FieldLogger, restClient *client.ClientImpl, summoner lol.Summoner) {
	logrusLogger.Debugf("retrieving match overview")

	var prompt string

	statusCode, err := restClient.Get("/match/overview").
		WithContext(context.Background()).
		WithJsonBody(handler.MatchOverviewRequest{SummonerPuuid: summoner.Puuid}).
		DoAndUnmarshal(&prompt)

	if err != nil {
		logrusLogger.WithError(err).Fatal("failed to get match overview")
		return
	}

	if statusCode >= 300 {
		err = fmt.Errorf("request failed with status code: %d", statusCode)
		logrusLogger.WithError(err).Fatal("failed to get match overview")
		return
	}

	printToStdOut(prompt)
}

func printToStdOut(msg string) {
	fmt.Println("================================================================================")
	fmt.Println(msg)
	fmt.Println("================================================================================")
}
