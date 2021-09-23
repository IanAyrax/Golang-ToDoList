package model

import "time"

type UserResponse struct{
	Id			int			`json:"id_user"`
	FullName	string		`json:"full_name"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
	RoleId		int			`json:"id_role"`
	CreatedAt	time.Time	`json:"created_at,omitempty"`
	UpdatedAt	time.Time	`json:"updated_at,omitempty"`
}