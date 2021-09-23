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

func NewUserRouter(router *httprouter.Router, db *sql.DB, validate *validator.Validate) *httprouter.Router {
	fmt.Println("User Router OK")
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router.GET("/api/user", userController.GetAll)
	router.GET("/api/user/:userId", userController.FindById)
	router.POST("/api/user", userController.Create)
	router.PUT("/api/user/:userId", userController.Update)
	router.DELETE("/api/user/:userId", userController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}