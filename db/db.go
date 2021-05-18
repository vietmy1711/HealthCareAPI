package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Sql struct {
	DB *sqlx.DB
	Host string
	Port int
	Username string
	Password string
	Dbname string
}

func (s *Sql) Connect() {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", s.Host, s.Port, s.Username, s.Password, s.Dbname)
	s.DB = sqlx.MustConnect("postgres", dataSource)
	if err := s.DB.Ping(); err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println("Connected to database....")
}

func (s *Sql) Close() {
	s.DB.Close()
}
