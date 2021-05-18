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
		Host : "ec2-52-71-161-140.compute-1.amazonaws.com",
		Port : 5432,
		Username: "gypzkqyxameflw",
		Password: "fa0bd299af4a929d7e232ea777cde4def55217b7f0e65ec698966c5b35052c72",
		Dbname: "d5e9m1htvn9vqg",
	}
	sql.Connect()
	defer sql.Close()
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	healthHandler := handler.HealthdayHandler{
		HealthdayRepo: repo_impl.NewHealthRepo(sql),
	}

	e := echo.New()
	api := router.API {
		Echo:       e,
		UserHandler: userHandler,
		HealthdayHandler: healthHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}



