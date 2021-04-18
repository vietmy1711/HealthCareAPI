package main

import (
	"github.com/heroku/go-getting-started/db"
	"github.com/heroku/go-getting-started/handler"
	"github.com/heroku/go-getting-started/log"
	"github.com/heroku/go-getting-started/repository/repo_impl"
	"github.com/heroku/go-getting-started/router"
	"github.com/labstack/echo/v4"
	"os"
	_ "os"
)

func init()  {
	os.Setenv("healthcare", "github")
	log.InitLogger(false)

}

func main() {
	sql := &db.Sql{
		Host : "localhost",
		Port : 5432,
		Username: "postgres",
		Password: "phucleuit",
		Dbname: "health_api",
	}
	sql.Connect()
	defer sql.Close()
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	e := echo.New()
	api := router.API {
		Echo:       e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}



