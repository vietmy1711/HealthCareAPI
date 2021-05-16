package repo_impl

import (
	"context"
	"github.com/heroku/go-getting-started/db"
	"github.com/heroku/go-getting-started/log"
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/repository"
	"github.com/lib/pq"
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
	_, err := u.sql.DB.NamedExecContext(context, statement, health)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return health, nil
			}
		}
		return health, nil
	}

	return health, nil
}

func (u HealthDayRepoImpl) GetInfoHealth(context context.Context, health model.HealthDay) (model.HealthDay, error) {
	var healthday = model.HealthDay{}
	statement := `
		INSERT INTO "user"(user_id, full_name, gender, blood)
		VALUES(:user_id, :full_name, :gender, :blood)
	`
	err := u.sql.DB.GetContext(context, &healthday, statement)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return healthday, nil
			}
		}
		return healthday, nil
	}

	return healthday, nil
}

func NewHealthRepo(sql *db.Sql) repository.HealthyRepo {
	return HealthDayRepoImpl{
		sql: sql,
	}
}

