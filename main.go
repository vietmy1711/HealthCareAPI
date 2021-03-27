package main

import (
	"github.com/heroku/go-getting-started/db"
	"github.com/heroku/go-getting-started/handler"
	"github.com/labstack/echo/v4"
)


func main() {

	sql := &db.Sql{
		Host: "localhost",
		Port: 5432,
		Username: "postgres",
		Password: "phucleuit",
		Dbname: "health_api",

	}
	sql.Connect()
	defer sql.Closed()
	e := echo.New()
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.Logger.Fatal(e.Start(":3000"))
}



