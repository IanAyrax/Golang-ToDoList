package service

import(
	"context"
	"example.com/GolangAPI2/model"
)

type UserService interface {
	Create(ctx context.Context, request model.UserCreateRequest) model.UserResponse
	Update(ctx context.Context, request model.UserUpdateRequest, roleId string, loggedId string) model.UserResponse
	Delete(ctx context.Context, roleId string, loggedId string, userId int)
	FindById(ctx context.Context, roleId string, loggedId string, userId int) model.UserResponse
	GetAll(ctx context.Context, roleId string) []model.UserResponse
}