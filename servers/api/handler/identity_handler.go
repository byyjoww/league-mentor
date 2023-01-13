//nolint
package handler

import (
	"net/http"

	"github.com/byyjoww/league-mentor/bll"
	app "github.com/byyjoww/league-mentor/services/http"
	"github.com/byyjoww/league-mentor/services/http/server"
	"github.com/byyjoww/league-mentor/services/http/server/response"
)

type IdentityHandler struct {
	decoder    server.Decoder
	controller bll.Controller
}

type IdentityRequest struct {
	SummonerName string `json:"summoner_name"`
}

func NewIdentityHandler(decoder server.Decoder, controller bll.Controller) *IdentityHandler {
	return &IdentityHandler{
		decoder:    decoder,
		controller: controller,
	}
}

func (h *IdentityHandler) GetMethod() string {
	return http.MethodGet
}

func (h *IdentityHandler) GetPath() string {
	return "/identity"
}

func (h *IdentityHandler) Handle(logger app.Logger, r *http.Request) server.Response {
	var (
		ctx = r.Context()
	)

	req := IdentityRequest{}
	if err := h.decoder.DecodeRequest(r, &req); err != nil {
		logger.WithError(err).Error("error decoding request")
		return response.NewJsonBadRequest(err)
	}

	logger = logger.WithFields(map[string]interface{}{
		"summoner_name": req.SummonerName,
	})

	identity, err := h.controller.GetIdentity(ctx, req.SummonerName)
	if err != nil {
		logger.WithError(err).Error("failed to get identity")
		return response.NewJsonInternalServerError(err)
	}

	logger.WithField("identity", identity).Info("successfully got identity")
	return response.NewJsonStatusOK(identity)
}
