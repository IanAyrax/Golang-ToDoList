package controller

import (
	"github.com/julienschmidt/httprouter"
	"example.com/GolangAPI2/service"
	"example.com/GolangAPI2/model"
	"example.com/GolangAPI2/helper"
	"net/http"
	"strconv"
	"fmt"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := model.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := model.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("User GetAll Controller OK")
	toDoResponses := controller.UserService.GetAll(request.Context())
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	toDoResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}