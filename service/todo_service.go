package service

import(
	"context"
	"example.com/GolangAPI2/model"
)

type ToDoService interface {
	Create(ctx context.Context, request model.ToDoCreateRequest) model.ToDoResponse
	Update(ctx context.Context, request model.ToDoUpdateRequest) model.ToDoResponse
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, todoId int) model.ToDoResponse
	GetAll(ctx context.Context) []model.ToDoResponse
}