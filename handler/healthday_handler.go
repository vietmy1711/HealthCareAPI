package handler

import (
	"github.com/heroku/go-getting-started/log"
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/model/req"
	"github.com/heroku/go-getting-started/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthdayHandler struct {
	HealthdayRepo repository.HealthyRepo
}

func (u *HealthdayHandler) HandleSaveHealthDay(c echo.Context) error {
	req := req.ReqHealthDay{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}
	healthday := model.HealthDay{
		Userid: req.Userid,
		Createat: req.Createat,
		Water: req.Water,
		Steps: req.Steps,
		Heartrate: req.Heartrate,
		Calogries: req.Calogries,
		Height: req.Height,
		Weight: req.Weight,
	}
	print(req.Userid)
	_, err := u.HealthdayRepo.SaveHealthDay(c.Request().Context(), healthday)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Xử lý thành công",
		Data: nil,
	})
}

func (u *HealthdayHandler) HandleGetHealthDay(c echo.Context) error {
	req := req.ReqHealthDay{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	healthday := model.HealthDay{

	}

	user, err := u.HealthdayRepo.GetInfoHealth(c.Request().Context(), healthday)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       user,
	})
}
