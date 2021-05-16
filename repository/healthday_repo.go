package repository

import (
	"context"
	"github.com/heroku/go-getting-started/model"
)

type HealthyRepo interface {
	SaveHealthDay(context context.Context, health model.HealthDay) (model.HealthDay, error)
	GetInfoHealth(context context.Context, health model.HealthDay) (model.HealthDay, error) // get 7 day about info health
}
