package model

import (
	"time"
)

type ToDo struct {
	Id	int
	UserId int
	Title string
	CreatedAt time.Time
	UpdatedAt time.Time
}