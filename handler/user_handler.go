package handler

import (
	"fmt"
	"github.com/heroku/go-getting-started/banana"
	"github.com/heroku/go-getting-started/log"
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/model/req"
	"github.com/heroku/go-getting-started/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSignUp{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user := model.User{
		UserId: req.Userid,
		FullName: req.FullName,
		Gender: req.Gender,
		Blood: req.Blood,
		Age: req.Age,
		Token: req.Token,
	}

	user, err := u.UserRepo.SaveUser(c.Request().Context(), user)
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
		Data:       nil,
	})
}

func (u *UserHandler) GetUser(c echo.Context) error {
	req := req.ReqGetUser{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user, err := u.UserRepo.GetUser(c.Request().Context(), req.Userid)
	fmt.Printf("gender:")
	fmt.Printf(string(user.Gender))
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "User Khong Ton Tai",
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       user,
	})
}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	req := req.ReqUpdateUser{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user := model.User{
		UserId: req.Userid,
		Token: req.Token,
	}
	user, err := u.UserRepo.Update(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    banana.UserNotFound.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}

func (u *UserHandler) UpdateUserInfo(c echo.Context) error {
	req := req.ReqUpdateInfo{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user := model.User{
		UserId: req.Userid,
		FullName: req.FullName,
		Age: req.Age,
	}
	user, err := u.UserRepo.UpdateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    banana.UserNotFound.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}
