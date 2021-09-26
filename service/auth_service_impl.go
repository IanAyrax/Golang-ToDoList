package service

import(
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"example.com/GolangAPI2/model"
	"example.com/GolangAPI2/repository"
	"example.com/GolangAPI2/helper"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
	//"example.com/GolangAPI2/exception"
	"fmt"
)

type AuthServiceImpl struct {
	AuthRepository 	repository.AuthRepository
	DB				*sql.DB
	Validate		*validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl {
		AuthRepository:	authRepository,
		DB:				DB,
		Validate:		validate,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request model.AuthLoginRequest) string {
	fmt.Println("Auth Service OK!")
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := model.User{
		Email:	request.Email,
		Password:	request.Password,
	}

	logged_user, err := service.AuthRepository.Login(ctx, tx, user)

	if err == nil {
		fmt.Println("Not Null")
		var err error
		//Creating Access Token
		os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
		atClaims := jwt.MapClaims{}
		atClaims["authorized"] = true
		atClaims["user_id"] = logged_user.UserId
		atClaims["role_id"] = logged_user.RoleId
		atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()	
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
		token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

		if err == nil {
			return token
		}else{
			return "Cannot Create Token"
		}
	}else{
		helper.PanicIfError(err)
		return "JWT Error"
	}
}

// func (service *AuthServiceImpl) Register(ctx context.Context) model.UserResponse {
// 	fmt.Println("Service OK")
// 	tx, err := service.DB.Begin()
// 	helper.PanicIfError(err)
// 	defer helper.CommitOrRollback(tx)

// 	users := service.UserRepository.GetAll(ctx, tx)

// 	return helper.ToUserResponses(users)
// }