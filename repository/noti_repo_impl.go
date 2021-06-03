package repository

import (
	"github.com/heroku/go-getting-started/model"
	"golang.org/x/net/context"
)

type NotiRepo interface {
	Checkin(context context.Context, userid string) (model.User, error)
}
