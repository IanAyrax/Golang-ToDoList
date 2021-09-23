package router

import(
	"database/sql"
	//"github.com/gofiber/fiber/v2"
	"github.com/julienschmidt/httprouter"
	"github.com/go-playground/validator/v10"
	"example.com/GolangAPI2/repository"
	"example.com/GolangAPI2/service"
	"example.com/GolangAPI2/controller"
	"example.com/GolangAPI2/exception"
	"fmt"
)

func NewToDoRouter(router *httprouter.Router, db *sql.DB, validate *validator.Validate) *httprouter.Router {
	fmt.Println("ToDo Router OK")
	todoRepository := repository.NewToDoRepository()
	todoService := service.NewToDoService(todoRepository, db, validate)
	todoController := controller.NewToDoController(todoService)

	router.GET("/api/todo", todoController.GetAll)
	router.GET("/api/todo/:todoId", todoController.FindById)
	router.POST("/api/todo", todoController.Create)
	router.PUT("/api/todo/:todoId", todoController.Update)
	router.DELETE("/api/todo/:todoId", todoController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}