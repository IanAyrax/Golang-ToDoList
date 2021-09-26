package model

type UserUpdateRequest struct{
	Id int `validate:"required"`
	FullName string `validate:"required,min=1,max=100" json:"full_name"`
	Email string `validate:"required,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
	RoleId int `validate:"required,numeric" json:"id_role"`
}