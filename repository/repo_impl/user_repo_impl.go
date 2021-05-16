package repo_impl

import (
	"context"
	"github.com/heroku/go-getting-started/db"
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/repository"
	"github.com/lib/pq"
	"github.com/heroku/go-getting-started/log"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
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
