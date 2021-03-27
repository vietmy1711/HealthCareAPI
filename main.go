package main

import (
	"github.com/heroku/go-getting-started/db"
	"github.com/labstack/echo/v4"
	"os"
)


func main() {
	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "7000"
	}
	DB := &db.Sql{}
	DB.Connect()
	defer DB.Close()

	e := echo.New()
	e.Logger.Fatal(e.Start(":" + port))
}



