package service

import(
	"context"
	"example.com/GolangAPI2/model"
)

type ToDoService interface {
	Create(ctx context.Context, request model.ToDoCreateRequest) model.ToDoResponse
	Update(ctx context.Context, get model.ToDoResponse, request model.ToDoUpdateRequest, roleId string, userId string) model.ToDoResponse
	Delete(ctx context.Context, get model.ToDoResponse, roleId string, userId string, todoId int)
	FindById(ctx context.Context, roleId string, userId string, todoId int) model.ToDoResponse
	GetAll(ctx context.Context, roleId string) []model.ToDoResponse
}