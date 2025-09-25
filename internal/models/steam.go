package models

type UserSummuryApiResponse struct {
	Response struct {
		Users []User `json:"players"`
	} `json:"response"`
}
