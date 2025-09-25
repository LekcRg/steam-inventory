package handlers

import (
	"fmt"
	"net/http"

	"github.com/LekcRg/steam-inventory/internal/api/middlewares"
	response "github.com/LekcRg/steam-inventory/internal/api/responder"
	"github.com/LekcRg/steam-inventory/internal/config"
	"github.com/LekcRg/steam-inventory/internal/service"
	"go.uber.org/zap"
)

type Handlers struct {
	log     *zap.Logger
	resp    *response.Responder
	config  *config.Config
	service *service.Service
}

func New(
	logger *zap.Logger, svc *service.Service,
	cfg *config.Config, resp *response.Responder,
) *Handlers {
	return &Handlers{
		log:     logger,
		resp:    resp,
		config:  cfg,
		service: svc,
	}
}

func (h *Handlers) Hi(w http.ResponseWriter, r *http.Request) {
	userIDAny := r.Context().Value(middlewares.CtxKeySteamID)
	userID, ok := userIDAny.(string)
	if !ok {
		h.resp.Error(w, http.StatusForbidden, "Forbidden")
		return
	}

	h.resp.Message(w, http.StatusOK,
		fmt.Sprintf("Hello user with Steam ID %s!", userID))
}
