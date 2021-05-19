package repo_impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/heroku/go-getting-started/banana"
	"github.com/heroku/go-getting-started/db"
	"github.com/heroku/go-getting-started/log"
	"github.com/heroku/go-getting-started/model"
	"github.com/lib/pq"
	"time"
	_ "time"
)

type HealthDayRepoImpl struct {
	sql *db.Sql
}

func (u HealthDayRepoImpl) SaveHealthDay(context context.Context, health model.HealthDay) (model.HealthDay, error) {
	statement := `
			INSERT INTO "healthday"(userid, createat, water, steps, heartrate, calogries, height, weight)
			VALUES(:userid, :createat, :water, :steps, :heartrate, :calogries, :height, :weight);
		`
	health.Createat = time.Now()
	_, err := u.sql.DB.NamedExecContext(context, statement, health)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return health, banana.HealthConflict
			}
		}
		return health, err
	}

	return health, nil
}

func (u HealthDayRepoImpl) GetInfoHealth(context context.Context, health string) ([]model.HealthDay, error) {
	//var healthday = model.HealthDay{}
	var listheathday []model.HealthDay
	var user model.User
	error := u.sql.DB.GetContext(context, &user, "SELECT * FROM account WHERE userid = $1", health)
	if error != nil {
		if error == sql.ErrNoRows {
			return listheathday, error
		}
		return listheathday, error
	}
	fmt.Printf("get health")
	err := u.sql.DB.SelectContext(context, &listheathday, "SELECT * FROM healthday WHERE userid = $1 ORDER BY createat ASC LIMIT $2", health, 7)
	if err != nil {
		log.Error(err.Error())
		return listheathday, err
	}
	return listheathday, nil
}

func NewHealthRepo(sql *db.Sql) HealthDayRepoImpl {
	return HealthDayRepoImpl{
		sql: sql,
	}
}

