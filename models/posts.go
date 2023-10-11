package models

import "time"

type Posts struct {
	Id        int
	UserId    int
	Title     string
	Content   string
	LikesCount int
	CreatedAt time.Time
	UpdatedAt time.Time
}