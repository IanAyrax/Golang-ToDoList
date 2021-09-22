package router

import(
//	"github.com/gofiber/fiber/v2"
	"github.com/julienschmidt/httprouter"
)

//func NewRouter() *fiber.App{
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	return router
}