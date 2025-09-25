package steam

import "fmt"

func (s *Steam) ApiURLUserSummary(steamID string) string {
	return fmt.Sprintf(
		"%s/ISteamUser/GetPlayerSummaries/v2/?key=%s&steamids=%s&format=json",
		s.config.Steam.APIDomain,
		s.config.Steam.APIKey,
		steamID,
	)
}
