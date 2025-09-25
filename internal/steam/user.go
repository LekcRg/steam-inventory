package steam

import (
	"fmt"

	"github.com/LekcRg/steam-inventory/internal/errs"
	"github.com/LekcRg/steam-inventory/internal/models"
)

func (s *Steam) GetUserSummary(steamID string) (*models.User, error) {
	user := &models.UserSummuryApiResponse{}
	apiURL := s.ApiURLUserSummary(steamID)

	res, err := s.client.R().
		SetResult(user).
		Get(apiURL)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("steam api error: %s", res.Status())
	}

	if len(user.Response.Users) == 0 {
		return nil, errs.ErrNotFoundUserSummary
	}

	return &user.Response.Users[0], nil
}
