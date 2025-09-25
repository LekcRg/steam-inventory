package steam

import (
	"github.com/LekcRg/steam-inventory/internal/config"
	"resty.dev/v3"
)

type Steam struct {
	client *resty.Client
	config *config.Config
}

func New(config *config.Config) *Steam {
	return &Steam{
		client: resty.New(),
		config: config,
	}
}
