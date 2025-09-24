package auth

import (
	"net/url"

	"github.com/LekcRg/steam-inventory/internal/api"
	"github.com/LekcRg/steam-inventory/internal/utils"
)

const steamURL = "https://steamcommunity.com/openid/login"

func GetRedirectURL(domain string) (*url.URL, error) {
	redirectURL, err := url.Parse(steamURL)
	if err != nil {
		return nil, err
	}

	returnToURL := domain + api.PathValidAuth

	queries := map[string]string{
		"openid.mode":       "checkid_setup",
		"openid.return_to":  returnToURL,
		"openid.realm":      domain,
		"openid.ns":         "http://specs.openid.net/auth/2.0",
		"openid.identity":   "http://specs.openid.net/auth/2.0/identifier_select",
		"openid.claimed_id": "http://specs.openid.net/auth/2.0/identifier_select",
	}

	redirectURL.RawQuery = utils.BuildQuery(queries)

	return redirectURL, nil
}

func GetValidURL(domain string, query url.Values) (*url.URL, error) {
	validURL, err := url.Parse(steamURL)
	if err != nil {
		return nil, err
	}

	query.Set("openid.mode", "check_authentication")
	validURL.RawQuery = query.Encode()

	return validURL, nil
}

// func Valid(domain string, query url.Values) (
// 	steamID64 string, isValid bool, err error,
// ) {
// 	validURL := GetValidURL(domain, query)

//
// }
