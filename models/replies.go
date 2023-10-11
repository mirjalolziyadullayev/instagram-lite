package models

import "time"

type Replies struct {
	Id         int
	UserId     int
	PostId     int
	CommentId  int
	Text       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
