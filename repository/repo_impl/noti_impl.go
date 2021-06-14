package repo_impl

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
	"github.com/heroku/go-getting-started/db"
	"github.com/heroku/go-getting-started/model"
	"golang.org/x/net/context"
)

const (
	serverKey = "AAAANwMoSnY:APA91bHFFcQV710bhsZV3M_zTLy85gcKvnFqRINKhltWtfCOubQ19BlBbljKTkQGU3aEA4Aiu4gf26nqImEQB9UFOzYUmCsZooZQ6kEyuTH3VVKtg3dalnYNO1etrQa_X7W2-abhshIR"
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

	data := map[string]string{
		"title": "Hello World1",
		"body": "Happy Day",
	}
	//message := fcm2.Notification{
	//	Title: "mon li",
	//	Body: "my mat lon",
	//}

	ids := []string{
		userid,
	}
	p := fcm.NewFcmClient(serverKey)
	p.NewFcmRegIdsMsg(ids, data)

	status, err := p.Send()
	user := model.User{}

	if err == nil {
		fmt.Printf("error");
		status.PrintResults()
		return user, err;
	} else {
		fmt.Printf("success");
		fmt.Println(err)
		return user, err;
	}
}

