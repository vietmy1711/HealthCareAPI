package handler

import (
	"github.com/heroku/go-getting-started/banana"
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
		Water: req.Water,
		Steps: req.Steps,
		Heartrate: req.Heartrate,
		Calories: req.Calories,
		Height: req.Height,
		Weight: req.Weight,
		ActiveEnergyBurned: req.ActiveEnergyBurned,
		BasalEnergyBurned: req.BasalEnergyBurned,
		BloodOxygen: req.BloodOxygen,
		DistanceWalkingRunning: req.DistanceWalkingRunning,
	}
	print(req.Userid)
	_, err := u.HealthdayRepo.SaveHealthDay(c.Request().Context(), healthday)
	if err != nil {
		if err == banana.HealthConflict {
			return c.JSON(http.StatusConflict, model.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
		}
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    banana.UserNotFound.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Xử lý thành công",
		Data: nil,
	})
}

func (u *HealthdayHandler) HandleGetInfoHealthInWeek(c echo.Context) error {
	req := req.ReqGetHealthDay{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	health, err := u.HealthdayRepo.GetInfoHealthInWeek(c.Request().Context(), req.Userid)
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
		Data:       health,
	})
}

func (u *HealthdayHandler) HandleGetInforHealthInDay(c echo.Context) error {
	req := req.ReqGetHealthDay{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	health, err := u.HealthdayRepo.GetInforHealthInDay(c.Request().Context(), req.Userid)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "User Khong Ton Tai",
			Data:       nil,
		})
	}
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
		Data:       health,
	})
}

func (u *HealthdayHandler) HandleGetWarning(c echo.Context) error {
	req := req.ReqGetHealthDay{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	health, err := u.HealthdayRepo.WarningHealth(c.Request().Context(), req.Userid)
	user, err := u.HealthdayRepo.GetUserForHeath(c.Request().Context(), req.Userid)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "User Khong Ton Tai",
			Data:       nil,
		})
	}
	firsthealthModel := model.HealthDay{}
	if health == nil {
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Khong co du lieu",
			Data:       nil,
		})
	} else {
		firsthealthModel = health[0]
	}

	var k int = 0
	var r float64 = 0
	if user.Gender == 0 {
		k = 5
	} else {
		k = -161
	}

	if health[0].Steps < 5000 {
		r = 1.2
	} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
		r = 1.55
	} else if health[0].Steps > 10000 {
		r = 1.725
	}

	var ststep, stwater, stkcal = 0,0,0 // 0 is bad, 1 is good, 2 is wonderfull
	if health[0].Steps < 5.000 {
		ststep = 0
	} else if health[0].Steps > 5000 && health[0].Steps <= 10000 {
		ststep = 1
	} else if health[0].Steps > 10000 {
		ststep = 2
	}

	if health[0].Water < 1000 {
		stwater = 0
	} else if health[0].Water > 1000 && health[0].Water <= 2000 {
		stwater = 1
	} else if health[0].Water > 2000 {
		stwater = 2
	}

	if user.Gender == 0 {
		if user.Age > 2 && user.Age <= 3 {
			if health[0].Steps < 5.000 {
				if health[0].Calories <= 1200 && health[0].Calories > 1000 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories <= 1400 && health[0].Calories > 1200 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories <= 1400 && health[0].Calories > 1200 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 4 && user.Age <= 8 {
			if health[0].Steps < 5.000 {
				if health[0].Calories <= 1400 && health[0].Calories > 1200 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories <= 1600 && health[0].Calories > 1400 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 2000 && health[0].Calories > 1600 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 9 && user.Age <= 13 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 2000 && health[0].Calories > 1600 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 2200 && health[0].Calories > 1800 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 2600 && health[0].Calories > 2000 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 14 && user.Age <= 18 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 2400 && health[0].Calories > 2000 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 2800 && health[0].Calories > 2400 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 3200 && health[0].Calories > 2800 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 19 && user.Age <= 30 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 2800 && health[0].Calories > 2400 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 3000 && health[0].Calories > 2800 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 3100 && health[0].Calories > 3000 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 31 && user.Age <= 50 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 2400 && health[0].Calories > 2000 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 2800 && health[0].Calories > 2400 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 3200 && health[0].Calories > 2800 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 51 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 2200 && health[0].Calories > 2000 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 2400 && health[0].Calories > 2200 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 2800 && health[0].Calories > 2400 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		}
	} else {
		if user.Age > 2 && user.Age <= 3 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 1000 && health[0].Calories > 900 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 1200 && health[0].Calories > 1000 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 1200 && health[0].Calories > 1000 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 4 && user.Age <= 8 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 1400 && health[0].Calories > 1200 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 1600 && health[0].Calories > 1400 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 1800 && health[0].Calories > 1400 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 9 && user.Age <= 13 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 1600 && health[0].Calories > 1400 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 2000 && health[0].Calories > 1600 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 2200 && health[0].Calories > 1800 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 14 && user.Age <= 18 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 1800 && health[0].Calories > 1700 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 2000 && health[0].Calories > 1900 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 2400 && health[0].Calories > 2300 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 19 && user.Age <= 30 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 2000 && health[0].Calories > 1800 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 2200 && health[0].Calories > 2000 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 2400 && health[0].Calories > 2300 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 31 && user.Age <= 50 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 1800 && health[0].Calories > 1700 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 2000 && health[0].Calories > 1900 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 2200 && health[0].Calories > 2100 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		} else if user.Age > 51 {
			if health[0].Steps < 5.000 {
				if health[0].Calories < 1600 && health[0].Calories > 1500 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 5000 && health[0].Steps < 10000 {
				if health[0].Calories < 1800 && health[0].Calories > 1700 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			} else if health[0].Steps > 10000 {
				if health[0].Calories < 2200 && health[0].Calories > 2000 {
					stkcal = 1
				} else {
					stkcal = 0
				}
			}
		}
	}



	warning := model.WaningModel{
		AboutKcal:  firsthealthModel.Calories,
		AboutStep:  firsthealthModel.Steps,
		AboutWater: firsthealthModel.Water,
		BRM:        (9.99 * firsthealthModel.Weight) + (6.25 * firsthealthModel.Height) - (4.92 * float64(user.Age)) - float64(k),
		TDEE:       ((9.99 * firsthealthModel.Weight) + (6.25 * firsthealthModel.Height) - (4.92 * float64(user.Age)) - float64(k)) * r,
		StatusKcal: stkcal,
		StatusStep: ststep,
		StatusWater: stwater,
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       warning,
	})
}

func (u *HealthdayHandler) HandleUpdateWater(c echo.Context) error {
	req := req.ReqWater{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	water := model.HealthDay{
		Userid: req.Userid,
		Water: req.Water,
	}
	result, err := u.HealthdayRepo.UpdateWater(c.Request().Context(), water)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Xử lý thành công",
		Data:       result,
	})
}

func (u *HealthdayHandler) HandleFakeHealthDay(c echo.Context) error {
	req := req.ReqFakeHealthDay{}
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
		Calories: req.Calories,
		Height: req.Height,
		Weight: req.Weight,
		ActiveEnergyBurned: req.ActiveEnergyBurned,
		BasalEnergyBurned: req.BasalEnergyBurned,
		BloodOxygen: req.BloodOxygen,
		DistanceWalkingRunning: req.DistanceWalkingRunning,
	}
	print(req.Userid)
	_, err := u.HealthdayRepo.FakeHealth(c.Request().Context(), healthday)
	if err != nil {
		if err == banana.HealthConflict {
			return c.JSON(http.StatusConflict, model.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
		}
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    banana.UserNotFound.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Xử lý thành công",
		Data: nil,
	})
}


