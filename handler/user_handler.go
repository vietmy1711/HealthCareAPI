package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"username" : "phuc le",
		"email" : "phuckhung58@gmail.com",
	})
}
