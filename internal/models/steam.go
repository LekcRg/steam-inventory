package models

type UserSummuryAPIResponse struct {
	Response struct {
		Users []User `json:"players"`
	} `json:"response"`
}
