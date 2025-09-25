package steam

import (
	"io"
	"net/url"
	"strings"

	"github.com/LekcRg/steam-inventory/internal/api"
	"github.com/LekcRg/steam-inventory/internal/errs"
	"github.com/LekcRg/steam-inventory/internal/querystring"
)

const steamURL = "https://steamcommunity.com/openid/login"

func (s *Steam) GetRedirectURL() (*url.URL, error) {
	redirectURL, err := url.Parse(steamURL)
	if err != nil {
		return nil, err
	}

	returnToURL := s.config.Domain + api.PathValidAuth

	queries := map[string]string{
		"openid.mode":       "checkid_setup",
		"openid.return_to":  returnToURL,
		"openid.realm":      s.config.Domain,
		"openid.ns":         "http://specs.openid.net/auth/2.0",
		"openid.identity":   "http://specs.openid.net/auth/2.0/identifier_select",
		"openid.claimed_id": "http://specs.openid.net/auth/2.0/identifier_select",
	}

	redirectURL.RawQuery = querystring.BuildQuery(queries)

	return redirectURL, nil
}

func (s *Steam) GetValidURL(query url.Values) (*url.URL, error) {
	validURL, err := url.Parse(steamURL)
	if err != nil {
		return nil, err
	}

	query.Set("openid.mode", "check_authentication")
	validURL.RawQuery = query.Encode()

	return validURL, nil
}

func (s *Steam) Valid(query url.Values) (
	steamID string, err error,
) {
	validURL, err := s.GetValidURL(query)
	if err != nil {
		return "", err
	}

	res, err := s.client.R().
		Post(validURL.String())
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.IsError() {
		return "", errs.ErrAuthValidationReq
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	isValid := strings.Contains(string(body), "is_valid:true")

	if !isValid {
		return "", errs.ErrInvalidAuth
	}

	claimedID := query.Get("openid.claimed_id")
	steamID = strings.TrimSpace(strings.Split(claimedID, "https://steamcommunity.com/openid/id/")[1])

	if steamID == "" {
		return "", errs.ErrInvalidSteamID
	}

	return steamID, nil
}
