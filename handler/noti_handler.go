package handler

import (
	"fmt"
	"github.com/heroku/go-getting-started/log"
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/model/req"
	"github.com/heroku/go-getting-started/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NotiHandler struct {
	NotiRepo repository.NotiRepo
}

func (u *NotiHandler) PushNoti(c echo.Context) error {
	req := req.NotiUser{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	fmt.Println("token")
	fmt.Printf(req.Token)
	_, err := u.NotiRepo.Checkin( c.Request().Context() ,req.Token)

	return err
}
