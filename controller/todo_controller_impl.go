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
	
	err := helper.VerifyToken(request)
	if err != nil {
		helper.PanicIfError(err)
	}

	fmt.Println(toDoCreateRequest.Title)
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

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	toDoUpdateRequest.Id = id

	toDoResponse := controller.ToDoService.Update(request.Context(), toDoUpdateRequest)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	toDoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	controller.ToDoService.Delete(request.Context(), id)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	toDoResponse := controller.ToDoService.FindById(request.Context(), id)
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	toDoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("GetAll Controller OK")
	toDoResponses := controller.ToDoService.GetAll(request.Context())
	webResponse := model.WebResponse{
		Code:	200,
		Status:	"OK",
		Data:	toDoResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}