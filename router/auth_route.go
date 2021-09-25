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

func NewAuthRouter(router *httprouter.Router, db *sql.DB, validate *validator.Validate) *httprouter.Router {
	fmt.Println("Auth Router OK")
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, db, validate)
	authController := controller.NewAuthController(authService)

	// router.POST("/api/register", authController.Register)
	router.POST("/api/login", authController.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}