package main

import (
	"example.com/GolangAPI2/router"
	"example.com/GolangAPI2/app"
	"example.com/GolangAPI2/helper"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-playground/validator/v10"
)

func main(){
	db := app.NewDB()
	validate := validator.New()
	mainRouter := router.NewRouter()

	mainRouter = router.NewToDoRouter(mainRouter, db, validate)

	server := http.Server{
		Addr:		app.GoDotEnvVariable("APP_HOST_DEV") + ":" + app.GoDotEnvVariable("APP_PORT"),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}