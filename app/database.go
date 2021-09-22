package app

import (
	"database/sql"
	"example.com/GolangAPI2/helper"
	"time"
)

func NewDB() *sql.DB {
	connectionString := GoDotEnvVariable("DB_USERNAME") + 
			":" + GoDotEnvVariable("DB_PASSWORD") +
			"@tcp(" + GoDotEnvVariable("DB_HOST") +
			":" + GoDotEnvVariable("DB_PORT") +
			")/" + GoDotEnvVariable("DB_DATABASE") +
			"?parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connectionString)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}