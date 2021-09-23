package service

import(
	"context"
	"example.com/GolangAPI2/model"
)

type UserService interface {
	Create(ctx context.Context, request model.UserCreateRequest) model.UserResponse
	Update(ctx context.Context, request model.UserUpdateRequest) model.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) model.UserResponse
	GetAll(ctx context.Context) []model.UserResponse
}