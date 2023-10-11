package models

import "time"

type Comments struct {
	Id        int
	UserId    int
	PostId    int
	Text      string
	CreatedAt time.Time
}