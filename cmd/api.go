package cmd

import (
	"github.com/byyjoww/league-mentor/bll"
	"github.com/byyjoww/league-mentor/config"
	"github.com/byyjoww/league-mentor/logging"
	apiApp "github.com/byyjoww/league-mentor/servers/api/http"
	"github.com/byyjoww/league-mentor/services/http/decoder"
	"github.com/byyjoww/league-mentor/services/http/server"
	"github.com/sirupsen/logrus"
)

func startApi(logrusLogger logrus.FieldLogger, cfg config.Config) {
	appLogger := logging.NewAppLogger(logrusLogger)
	decoder := decoder.New()

	controller := bll.NewGenericController(cfg)
	api := apiApp.New(appLogger, decoder, cfg.Http.Api, controller)

	logrus.Info("App initialized succesfully")
	if err := server.ListenAndServe(api); err != nil {
		logrus.WithError(err).Fatal("failed to start server")
	}
}
