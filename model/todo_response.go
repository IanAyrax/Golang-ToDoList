package model

import "time"

type ToDoResponse struct{
	Id	int	`json:"id_todo"`
	UserId	int	`json:"user_id"`
	Title	string	`json:"title"`
	CreatedAt	time.Time	`json:"created_at,omitempty"`
	UpdatedAt	time.Time	`json:"updated_at,omitempty"`
}