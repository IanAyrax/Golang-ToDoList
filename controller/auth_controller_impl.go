package controller

import (
	"github.com/julienschmidt/httprouter"
	"example.com/GolangAPI2/service"
	"example.com/GolangAPI2/model"
	"example.com/GolangAPI2/helper"
	"net/http"
	"fmt"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Auth Controller OK!")
	authLoginRequest := model.AuthLoginRequest{}
	helper.ReadFromRequestBody(request, &authLoginRequest)
	fmt.Println(authLoginRequest.Email)

	loginResponse := controller.AuthService.Login(request.Context(), authLoginRequest)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	loginResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// fmt.Println("User GetAll Controller OK")
	// toDoResponses := controller.UserService.GetAll(request.Context())
	// webResponse := model.WebResponse{
	// 	Code:	200,
	// 	Status:	"OK",
	// 	Data:	toDoResponses,
	// }

	// helper.WriteToResponseBody(writer, webResponse)
}