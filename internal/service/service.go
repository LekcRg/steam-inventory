package service

import (
	"context"

	"github.com/LekcRg/steam-inventory/internal/config"
	"github.com/LekcRg/steam-inventory/internal/models"
	"github.com/LekcRg/steam-inventory/internal/steam"
)

type repository interface {
	CreateOrUpdateUser(ctx context.Context, user *models.User) (*models.User, error)
}

type Service struct {
	repo   repository
	config *config.Config
	steam  *steam.Steam
}

func New(
	config *config.Config, repo repository, st *steam.Steam,
) *Service {
	return &Service{
		repo:   repo,
		config: config,
		steam:  st,
	}
}
