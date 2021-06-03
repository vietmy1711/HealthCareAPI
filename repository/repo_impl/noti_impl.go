package repo_impl

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
	fcm2 "github.com/appleboy/go-fcm"
	"github.com/heroku/go-getting-started/db"
	"github.com/heroku/go-getting-started/model"
	"golang.org/x/net/context"
)

const (
	serverKey = "AAAANwMoSnY:APA91bFJ9gh3wO5SvMkZjIZoVXfoMEF5Ced4IIMJAi5JE_ILYuhq1BKGNY6o10M67xckvC9WkIlY57rv-NsIgg3hiOyaPmyefa1Ls_vOa5z1WbkLgk6IZ7bRHixA6mAEKTcpR9bcAZIC"
)

type NotiRepoImpl struct {
	sql *db.Sql
}

func NewNotiRepo(sql *db.Sql) NotiRepoImpl {
	return NotiRepoImpl{
		sql: sql,
	}
}

func (u NotiRepoImpl) Checkin(context context.Context, userid string) (model.User, error) {
	//data := map[string]string{
	//	"title": "Hello World1",
	//	"body": "Happy Day",
	//}

	message := fcm2.Notification{
		Title: "mon li",
		Body: "my mat lon",
	}

	ids := []string{
		userid,
	}
	p := fcm.NewFcmClient(serverKey)
	p.NewFcmRegIdsMsg(ids, message)
	//p.AppendDevices(xds)

	status, err := p.Send()
	user := model.User{}

	if err == nil {
		status.PrintResults()
		return user, err;
	} else {
		fmt.Println(err)
		return user, err;
	}
}

