package models

import "time"

type Comments struct {
	CreatedAt time.Time
	User      User
	Text      string
}
