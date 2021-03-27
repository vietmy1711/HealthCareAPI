package db

import (
	"fmt"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type Sql struct {
	Db *sqlx.DB
	Host string
	Port int
	Username string
	Password string
	Dbname string
}

func (s *Sql) Connect() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.Username, s.Password, s.Dbname)
	s.Db = sqlx.MustConnect("postgres", datasource)
	if err := s.Db.Ping(); err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println("Connect database")
}

func (s *Sql) Closed() {
	s.Db.Close()
}
