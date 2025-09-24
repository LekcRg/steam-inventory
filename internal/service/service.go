package service

import (
	"github.com/LekcRg/steam-inventory/internal/config"
)

type repository interface {
	GetUsersIds() ([]int, error)
}

type Service struct {
	repo   repository
	config *config.Config
}

func New(config *config.Config, repo repository) *Service {
	return &Service{
		repo:   repo,
		config: config,
	}
}
