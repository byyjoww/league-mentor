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
	tickInterval = 20 * time.Second
)

func startClient(logrusLogger logrus.FieldLogger, cfg config.Config) {
	var (
		summoner lol.Summoner
	)

	restClient := client.New(http.DefaultClient, decoder.New()).
		WithAuth(auth.NewBasic("user", "pass")).
		WithBaseUrl("http://0.0.0.0:8080" + "/api/v1")

	statusCode, err := restClient.Get("/identity").
		WithContext(context.Background()).
		WithJsonBody(handler.IdentityRequest{SummonerName: cfg.RiotGames.SummonerName}).
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
			getLaneAdvice(logger, restClient, summoner)
		}
	}
}

func getLaneAdvice(logrusLogger logrus.FieldLogger, restClient *client.ClientImpl, summoner lol.Summoner) {
	logrusLogger.Debugf("retrieving lane advice")

	var prompt string

	statusCode, err := restClient.Get("/advice/lane").
		WithContext(context.Background()).
		WithJsonBody(handler.LaneAdviceRequest{SummonerPuuid: summoner.Puuid}).
		DoAndUnmarshal(&prompt)

	if err != nil {
		logrusLogger.WithError(err).Fatal("failed to get lane advice")
		return
	}

	if statusCode >= 300 {
		err = fmt.Errorf("request failed with status code: %d", statusCode)
		logrusLogger.WithError(err).Fatal("failed to get lane advice")
		return
	}

	printToStdOut(prompt)
}

func printToStdOut(msg string) {
	fmt.Println("================================================================================")
	fmt.Println(msg)
	fmt.Println("================================================================================")
}
