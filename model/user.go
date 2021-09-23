package model

import (
	"time"
)

type User struct {
	UserId	int
	FullName string
	Email string
	Password string
	RoleId int
	CreatedAt time.Time
	UpdatedAt time.Time
}