package main

import (
	"example.com/GolangAPI2/router"
	"example.com/GolangAPI2/app"
	"example.com/GolangAPI2/helper"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-playground/validator/v10"
	"fmt"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

func main(){
	db := app.NewDB()
	validate := validator.New()
	mainRouter := router.NewRouter()

	mainRouter = router.NewToDoRouter(mainRouter, db, validate)
	mainRouter = router.NewUserRouter(mainRouter, db, validate)
	mainRouter = router.NewAuthRouter(mainRouter, db, validate)

	server := http.Server{
		Addr:		app.GoDotEnvVariable("APP_HOST_DEV") + ":" + app.GoDotEnvVariable("APP_PORT"),
		Handler:	&Logger{mainRouter},
	}
	
	fmt.Println("Serve :" + app.GoDotEnvVariable("APP_HOST_DEV") + ":" + app.GoDotEnvVariable("APP_PORT"))
	err := server.ListenAndServe()
	fmt.Println("Listening")
	helper.PanicIfError(err)
}