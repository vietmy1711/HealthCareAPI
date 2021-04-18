package repository

import (
	"context"
	"github.com/heroku/go-getting-started/model"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
}
