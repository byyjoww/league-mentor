package http

import (
	"net/http"

	"github.com/byyjoww/league-mentor/bll"
	"github.com/byyjoww/league-mentor/config"
	"github.com/byyjoww/league-mentor/servers/api/handler"
	app "github.com/byyjoww/league-mentor/services/http"
	"github.com/byyjoww/league-mentor/services/http/server"
	appMiddleware "github.com/byyjoww/league-mentor/services/http/server/middleware"
)

func New(logger app.Logger, decoder server.Decoder, appConfig config.HTTPServer, controller bll.Controller) server.App {
	basicAuth := appMiddleware.NewBasicAuthMiddleware(
		appConfig.Auth.User,
		appConfig.Auth.Pass,
		appConfig.Auth.Enabled,
	)

	panicLogger := appMiddleware.NewPanicLoggerMiddleware()

	mux := server.NewMux(logger).
		WithMiddlewares(basicAuth, panicLogger).
		PathPrefixSubrouter("/api/v1")

	mux.AddHandlers(
		handler.NewLaneAdviceHandler(decoder, controller),
		handler.NewIdentityHandler(decoder, controller),
	)

	return &http.Server{
		Addr:    appConfig.Address,
		Handler: mux,
	}
}
