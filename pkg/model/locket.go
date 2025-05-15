package model

import "time"

type LocketType string

var ValidLocketTypes = [3]LocketType{"image/jpeg", "image/png", "video/mp4"}

type Locket struct {
	ID        uint       `json:"id"`
	Type      LocketType `json:"type"`
	ImageUrl  string     `json:"imageUrl"`
	Caption   string     `json:"caption"`
	Username  string     `json:"username"`
	CreatedAt time.Time  `json:"createdAt"`
}
