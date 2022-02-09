package models

import "time"

type Comments struct {
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	User      User      `json:"user_id"`
}
