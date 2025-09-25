package service

import (
	"context"

	"github.com/LekcRg/steam-inventory/internal/cache"
	"github.com/LekcRg/steam-inventory/internal/config"
	"github.com/LekcRg/steam-inventory/internal/models"
	"github.com/LekcRg/steam-inventory/internal/steam"
)

type repository interface {
	CreateOrUpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserBySteamID(ctx context.Context, userID string) (*models.User, error)
}

type Service struct {
	repo   repository
	config *config.Config
	steam  *steam.Steam
	cache  *cache.Cache
}

func New(
	cfg *config.Config, repo repository,
	st *steam.Steam, c *cache.Cache,
) *Service {
	return &Service{
		repo:   repo,
		config: cfg,
		steam:  st,
		cache:  c,
	}
}
