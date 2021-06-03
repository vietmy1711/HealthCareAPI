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
			INSERT INTO "healthday"
			VALUES(:userid, :createat, :water, :steps, :heartrate, :calories, :height, :weight, :active_energy_burned, :basal_energy_burned, :blood_oxygen);
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

func (u HealthDayRepoImpl) GetInfoHealthInWeek(context context.Context, health string) ([]model.HealthDay, error) {
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
	fmt.Printf(health)
	err := u.sql.DB.SelectContext(context, &listheathday, "SELECT createat::DATE, MAX(WATER) AS WATER, MAX(STEPS) AS " +
		"STEPS, AVG(HEARTRATE) AS HEARTRATE, " +
		"AVG(CALORIES) AS CALORIES, MAX(HEIGHT) AS HEIGHT, MAX(WEIGHT) AS WEIGHT, " +
		"AVG(active_energy_bunred) AS active_energy_burned, AVG(basal_energy_bunred) " +
		"AS basal_energy_burned, AVG(blood_oxygen) AS blood_oxygen FROM healthday " +
		"WHERE userid = $1 AND createat > CURRENT_DATE - 7 " +
		"GROUP BY createat::DATE ORDER BY createat::DATE DESC", health)
	if err != nil {
		log.Error(err.Error())
		return listheathday, err
	}
	return listheathday, nil
}

func (u HealthDayRepoImpl) GetInforHealthInDay(context context.Context, userid string) ([]model.HealthDay, error) {
	//var healthday = model.HealthDay{}
	var listheathday []model.HealthDay
	var user model.User
	error := u.sql.DB.GetContext(context, &user, "SELECT * FROM account WHERE userid = $1", userid)
	if error != nil {
		if error == sql.ErrNoRows {
			return listheathday, error
		}
		return listheathday, error
	}
	fmt.Printf("get health day")
	print(time.Now().String())
	err := u.sql.DB.SelectContext(context, &listheathday, "SELECT water, steps, heartrate, calories, height, weight, active_energy_bunred as active_energy_burned," +
		"basal_energy_bunred as basal_energy_burned, blood_oxygen FROM healthday  WHERE createat::date = $1 AND userid = $2", time.Now(), userid )
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

