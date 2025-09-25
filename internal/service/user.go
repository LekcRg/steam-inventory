package service

import (
	"context"

	"github.com/LekcRg/steam-inventory/internal/models"
)

func (s *Service) UserInfo(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.repo.GetUserBySteamID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
