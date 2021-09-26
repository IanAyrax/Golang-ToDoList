package model

type UserCreateRequest struct{
	FullName string `validate:"required,min=1,max=100" json:"full_name"`
	Email string `validate:"required,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
	RoleId int `validate:"numeric" json:"id_role"`
}