package models

import "time"

type Comments struct {
	Id        int
	UserId    User
	PostId    Posts
	Text      string
	CreatedAt time.Time
}