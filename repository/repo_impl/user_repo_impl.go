package repo_impl

import (
	"context"
	"github.com/heroku/go-getting-started/db"
	"github.com/heroku/go-getting-started/log"
	"github.com/heroku/go-getting-started/model"
	"github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) UserRepoImpl {
	return UserRepoImpl{
		sql: sql,
	}
}

func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO "account"(userid, username, blood, gender)
		VALUES(:userid, :username, :blood, :gender)
	`
	_, err := u.sql.DB.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, nil
			}
		}
		return user, nil
	}

	return user, nil
}

func (u UserRepoImpl) GetUser(context context.Context, userid string) (model.User, error) {
	var user model.User
	err := u.sql.DB.GetContext(context, &user, "SELECT * FROM account WHERE userid = $1", userid)
	if err != nil {
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}


