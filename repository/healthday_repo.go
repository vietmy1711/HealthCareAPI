package repository

import (
	"context"
	"github.com/heroku/go-getting-started/model"
)

type HealthyRepo interface {
	SaveHealthDay(context context.Context, health model.HealthDay) (model.HealthDay, error)
	FakeHealth(context context.Context, health model.HealthDay) (model.HealthDay, error)
	GetInfoHealthInWeek(context context.Context, health string) ([]model.HealthDay, error) // get 7 day about info health
	GetInforHealthInDay(context context.Context, userid string) ([]model.HealthDay, error)
	UpdateWater(context context.Context, health model.HealthDay) (model.HealthDay, error)
	WarningHealth(context context.Context, userid string) ([]model.HealthDay, error)
	GetUserForHeath(context context.Context, userid string) (model.User, error)
}
