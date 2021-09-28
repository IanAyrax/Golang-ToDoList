package model

type ToDoCreateRequest struct{
	UserId int `json:"user_id"`
	Title string `validate:"required,min=1,max=100" json:"title"`
}