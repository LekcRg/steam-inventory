package middlewares

import (
	"github.com/LekcRg/steam-inventory/internal/config"
	"go.uber.org/zap"
)

type Middlewares struct {
	log    *zap.Logger
	config *config.Config
}

func New(cfg *config.Config, log *zap.Logger) *Middlewares {
	return &Middlewares{
		log:    log,
		config: cfg,
	}
}
