package middlewares

import (
	response "github.com/LekcRg/steam-inventory/internal/api/responder"
	"github.com/LekcRg/steam-inventory/internal/cache"
	"github.com/LekcRg/steam-inventory/internal/config"
	"go.uber.org/zap"
)

type Middlewares struct {
	log    *zap.Logger
	config *config.Config
	cache  *cache.Cache
	resp   *response.Responder
}

func New(
	cfg *config.Config, log *zap.Logger,
	c *cache.Cache, r *response.Responder,
) *Middlewares {
	return &Middlewares{
		log:    log,
		config: cfg,
		cache:  c,
		resp:   r,
	}
}
