package controller

import (
	"github.com/julienschmidt/httprouter"
	"example.com/GolangAPI2/service"
	"example.com/GolangAPI2/model"
	"example.com/GolangAPI2/helper"
	"example.com/GolangAPI2/middleware"
	"net/http"
	"strconv"
	"fmt"
)

type ToDoControllerImpl struct {
	ToDoService service.ToDoService
}

func NewToDoController(todoService service.ToDoService) ToDoController {
	return &ToDoControllerImpl{
		ToDoService: todoService,
	}
}

func (controller *ToDoControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("ToDoController Ok!")
	toDoCreateRequest := model.ToDoCreateRequest{}
	helper.ReadFromRequestBody(request, &toDoCreateRequest)
	
	userId, _, err := middleware.VerifyToken(request)
	helper.PanicIfError(err)
	
	i, _ := strconv.Atoi(userId) 
	toDoCreateRequest.UserId = i
	toDoResponse := controller.ToDoService.Create(request.Context(), toDoCreateRequest)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	toDoResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	toDoUpdateRequest := model.ToDoUpdateRequest{}
	helper.ReadFromRequestBody(request, &toDoUpdateRequest)
	userId, roleId, err := middleware.VerifyToken(request)
	request.Header.Set("RoleId", roleId)
	helper.PanicIfError(err)
	

	//err = helper.IsAdmin(request)
	//helper.PanicIfError(err)
	
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)
	
	toDoResponse := controller.ToDoService.FindById(request.Context(),roleId, userId, id)
	//if fmt.Sprintf("%v", toDoResponse.UserId) != userId || helper.IsAdmin(request) != nil{
	//	helper.PanicIfError(errors.New("Action Not Allowed : Not the Owner!!!!"))
	//}
	
	toDoUpdateRequest.Id = id

	toDoResponse = controller.ToDoService.Update(request.Context(), toDoResponse, toDoUpdateRequest, roleId, userId)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	toDoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId, roleId, err := middleware.VerifyToken(request)
	request.Header.Set("RoleId", roleId)
	helper.PanicIfError(err)
	

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	toDoResponse := controller.ToDoService.FindById(request.Context(), roleId, userId, id)
	//if fmt.Sprintf("%v", toDoResponse.UserId) != userId || helper.IsAdmin(request) != nil{
	//	helper.PanicIfError(errors.New("Action Not Allowed : Not the Owner!!!!"))
	//}

	controller.ToDoService.Delete(request.Context(), toDoResponse, roleId, userId, id)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId, roleId, err := middleware.VerifyToken(request)
	request.Header.Set("RoleId", roleId)
	if err != nil {
		helper.PanicIfError(err)
	}

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	toDoResponse := controller.ToDoService.FindById(request.Context(), roleId, userId, id)
	// if fmt.Sprintf("%v", toDoResponse.UserId) != userId || helper.IsAdmin(roleId) != nil{
	// 	helper.PanicIfError(errors.New("Action Not Allowed : Not the Owner!!!!"))
	// }

	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	toDoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("GetAll Controller OK")
	_, roleId, err := middleware.VerifyToken(request)
	request.Header.Set("RoleId", roleId)
	helper.PanicIfError(err)
	
	// err = helper.IsAdmin(roleId)
	// helper.PanicIfError(err)
	
	toDoResponses := controller.ToDoService.GetAll(request.Context(), roleId)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	toDoResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}