package helper

import (
	"example.com/GolangAPI2/model"
)

func ToUserResponse(value model.User) model.UserResponse {
	return model.UserResponse{
		Id:	value.UserId,
		FullName:	value.FullName,
		Email:	value.Email,
		Password: value.Password,
		RoleId:	value.RoleId,
		CreatedAt: value.CreatedAt,
		UpdatedAt: value.UpdatedAt,
	}
}

func ToUserResponses(users []model.User) []model.UserResponse{
	var userResponses []model.UserResponse
	for _, user := range users{
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}