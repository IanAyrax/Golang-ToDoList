package helper

import (
	"example.com/GolangAPI2/model"
)

func ToToDoResponse(value model.ToDo) model.ToDoResponse {
	return model.ToDoResponse{
		Id:	value.Id,
		UserId:	value.UserId,
		Title:	value.Title,
		CreatedAt: value.CreatedAt,
		UpdatedAt: value.UpdatedAt,
	}
}

func ToToDoResponses(todos []model.ToDo) []model.ToDoResponse{
	var todoResponses []model.ToDoResponse
	for _, todo := range todos{
		todoResponses = append(todoResponses, ToToDoResponse(todo))
	}

	return todoResponses
}