package models

import "time"

type User struct {
	ID                       int       `json:"id" db:"id"`
	SteamID                  string    `json:"steamid" db:"steamid"`
	CommunityVisibilityState int       `json:"communityvisibilitystate" db:"communityvisibilitystate"`
	PersonaName              string    `json:"personaname" db:"personaname"`
	Avatar                   string    `json:"avatarfull" db:"avatar"`
	LastLogOffSteam          int       `json:"lastlogoff" db:"lastlogoff_steam"`
	RealName                 string    `json:"realname" db:"realname"`
	TimeCreatedSteam         int       `json:"timecreated" db:"timecreated_steam"`
	CreatedAt                time.Time `json:"created_at" db:"created_at"`
	UpdateAt                 time.Time `json:"updated_at" db:"updated_at"`
}

type UserResponse struct {
	ID                       int    `json:"id" db:"id"`
	SteamID                  string `json:"steamid" db:"steamid"`
	PersonaName              string `json:"personaname" db:"personaname"`
	Avatar                   string `json:"avatar_full" db:"avatar"`
	RealName                 string `json:"realname" db:"realname"`
	CommunityVisibilityState int    `json:"communityvisibilitystate" db:"communityvisibilitystate"`
}
