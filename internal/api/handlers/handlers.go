package handlers

import (
	"net/http"

	response "github.com/LekcRg/steam-inventory/internal/api/responder"
	"github.com/LekcRg/steam-inventory/internal/config"
	"github.com/LekcRg/steam-inventory/internal/service"
	"go.uber.org/zap"
)

type Handlers struct {
	log     *zap.Logger
	resp    *response.Responder
	cfg     *config.Config
	service *service.Service
}

func New(
	logger *zap.Logger, svc *service.Service,
	cfg *config.Config, resp *response.Responder,
) *Handlers {
	return &Handlers{
		log:     logger,
		resp:    resp,
		cfg:     cfg,
		service: svc,
	}
}

func (h *Handlers) Hi(w http.ResponseWriter, r *http.Request) {
	h.resp.Message(w, http.StatusOK, "Hello, world!")
}
