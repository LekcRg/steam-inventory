package service

import (
	"context"
	"net/url"

	"github.com/LekcRg/steam-inventory/internal/crypto"
	"github.com/LekcRg/steam-inventory/internal/models"
)

func (s *Service) GetAuthRedirectURL() (*url.URL, error) {
	return s.steam.GetRedirectURL()
}

func (s *Service) AuthValid(
	ctx context.Context, query url.Values,
) (user *models.User, session string, err error) {
	steamID, err := s.steam.Valid(query)
	if err != nil {
		return nil, "", err
	}

	user, err = s.steam.GetUserSummary(steamID)
	if err != nil {
		return nil, "", err
	}

	_, err = s.repo.CreateOrUpdateUser(ctx, user)
	if err != nil {
		return nil, "", err
	}

	ses, err := crypto.GenSession()
	if err != nil {
		return nil, "", err
	}

	err = s.cache.SetSession(ctx, ses, steamID)
	if err != nil {
		return nil, "", err
	}

	return user, ses, err
}
