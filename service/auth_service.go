package service

import(
	"context"
	"example.com/GolangAPI2/model"
)

type AuthService interface {
	//Register(ctx context.Context, request model.UserCreateRequest) model.UserResponse
	Login(ctx context.Context, request model.AuthLoginRequest) string
}