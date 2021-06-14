package router

import (
	"github.com/heroku/go-getting-started/handler"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo *echo.Echo
	UserHandler handler.UserHandler
	HealthdayHandler handler.HealthdayHandler
	NotiHander handler.NotiHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
	api.Echo.POST("/health/save-health", api.HealthdayHandler.HandleSaveHealthDay)
	api.Echo.POST("/health/fake-health", api.HealthdayHandler.HandleFakeHealthDay)
	api.Echo.POST("/health/update-user", api.UserHandler.UpdateUser)
	api.Echo.POST("/health/checkin", api.NotiHander.PushNoti)
	api.Echo.POST("/health/update-water", api.HealthdayHandler.HandleUpdateWater)
	api.Echo.POST("/health/update-info", api.UserHandler.UpdateUserInfo)
	api.Echo.POST("/health/warning", api.HealthdayHandler.HandleGetWarning)

	api.Echo.POST("/health/get-user", api.UserHandler.GetUser)
	api.Echo.POST("/health/get-healthinweek", api.HealthdayHandler.HandleGetInfoHealthInWeek)
	api.Echo.POST("/health/get-healthinday", api.HealthdayHandler.HandleGetInforHealthInDay)

}
