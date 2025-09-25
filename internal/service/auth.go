package service

import (
	"context"
	"net/url"

	"github.com/LekcRg/steam-inventory/internal/models"
)

func (s *Service) GetAuthRedirectURL() (*url.URL, error) {
	return s.steam.GetRedirectURL()
}

func (s *Service) AuthValid(ctx context.Context, query url.Values) (*models.User, error) {
	steamID, err := s.steam.Valid(query)
	if err != nil {
		return nil, err
	}

	user, err := s.steam.GetUserSummary(steamID)
	if err != nil {
		return nil, err
	}

	user, err = s.repo.CreateOrUpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, err
}
