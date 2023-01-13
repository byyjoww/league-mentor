//nolint
package handler

import (
	"net/http"

	"github.com/byyjoww/league-mentor/bll"
	app "github.com/byyjoww/league-mentor/services/http"
	"github.com/byyjoww/league-mentor/services/http/server"
	"github.com/byyjoww/league-mentor/services/http/server/response"
)

type MatchOverviewHandler struct {
	decoder    server.Decoder
	controller bll.Controller
}

type MatchOverviewRequest struct {
	SummonerPuuid string `json:"summoner_puuid"`
}

func NewMatchOverviewHandler(decoder server.Decoder, controller bll.Controller) *MatchOverviewHandler {
	return &MatchOverviewHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *MatchOverviewHandler) GetMethod() string {
	return http.MethodGet
}

func (h *MatchOverviewHandler) GetPath() string {
	return "/match/overview"
}

func (h *MatchOverviewHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	var (
		ctx = r.Context()
	)

	req := MatchOverviewRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return response.NewJsonBadRequest(err)
	}

	logger = logger.WithFields(map[string]interface{}{
		"summoner_puuid": req.SummonerPuuid,
	})

	overview, err := h.controller.GetCurrentMatchOverview(ctx, req.SummonerPuuid)
	if err != nil {
		logger.WithError(err).Error("failed to get match overview")
		return response.NewJsonInternalServerError(err)
	}

	logger.WithField("overview", overview).Info("successfully got match overview")
	return response.NewJsonStatusOK(overview)
}
