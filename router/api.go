package router

import (
	"github.com/heroku/go-getting-started/handler"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo *echo.Echo
	UserHandler handler.UserHandler
	HealthdayHandler handler.HealthdayHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
	api.Echo.POST("/health/get-health", api.HealthdayHandler.HandleSaveHealthDay)
}
