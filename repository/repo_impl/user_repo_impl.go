package repo_impl

import (
	"context"
	"database/sql"
	"github.com/heroku/go-getting-started/banana"
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
		INSERT INTO "account"
		VALUES(:userid, :username, :blood, :gender, :age, :token)
	`
	_, err := u.sql.DB.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			log.Error(err.Code.Name())
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
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
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}
	return user, nil
}

func (u UserRepoImpl) Update(context context.Context, userid model.User) (model.User, error) {
	var user model.User
	err := u.sql.DB.GetContext(context, &user, "SELECT * FROM account WHERE userid = $1", userid.UserId)
	if err != nil {
		log.Error(err.Error())
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}
	sqlStatement := `
		UPDATE account SET token = :token WHERE userid = :userid
	`

	result, err := u.sql.DB.NamedExecContext(context, sqlStatement, userid)

	if err != nil {
		log.Error(err.Error())
		return user, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return user, banana.UserNotUpdated
	}
	if count == 0 {
		return user, banana.UserNotUpdated
	}
	return user, nil
}


