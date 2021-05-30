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
	api.Echo.POST("/health/save-health", api.HealthdayHandler.HandleSaveHealthDay)

	api.Echo.GET("/health/get-user", api.UserHandler.GetUser)
	api.Echo.GET("/health/get-healthinweek", api.HealthdayHandler.HandleGetInfoHealthInWeek)
	api.Echo.GET("/health/get-healthinday", api.HealthdayHandler.HandleGetInforHealthInDay)
}
