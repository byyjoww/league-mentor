//nolint
package handler

import (
	"net/http"

	"github.com/byyjoww/league-mentor/bll"
	app "github.com/byyjoww/league-mentor/services/http"
	"github.com/byyjoww/league-mentor/services/http/server"
	"github.com/byyjoww/league-mentor/services/http/server/response"
)

type LaneAdviceHandler struct {
	decoder    server.Decoder
	controller bll.Controller
}

type LaneAdviceRequest struct {
	SummonerPuuid string `json:"summoner_puuid"`
}

func NewLaneAdviceHandler(decoder server.Decoder, controller bll.Controller) *LaneAdviceHandler {
	return &LaneAdviceHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *LaneAdviceHandler) GetMethod() string {
	return http.MethodGet
}

func (h *LaneAdviceHandler) GetPath() string {
	return "/advice/lane"
}

func (h *LaneAdviceHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	var (
		ctx = r.Context()
	)

	req := LaneAdviceRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return response.NewJsonBadRequest(err)
	}

	logger = logger.WithFields(map[string]interface{}{
		"summoner_puuid": req.SummonerPuuid,
	})

	advice, err := h.controller.GetLaneAdvice(ctx, req.SummonerPuuid)
	if err != nil {
		logger.WithError(err).Error("failed to get lane advice")
		return response.NewJsonInternalServerError(err)
	}

	logger.WithField("advice", advice).Info("successfully got lane advice")
	return response.NewJsonStatusOK(advice)
}
