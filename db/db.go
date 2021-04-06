package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Sql struct {
	DB *sqlx.DB
}

func (s *Sql) Connect() {
	s.DB = sqlx.MustConnect("postgres", "host=ec2-52-71-161-140.compute-1.amazonaws.com port=5432 user=gypzkqyxameflw password=fa0bd299af4a929d7e232ea777cde4def55217b7f0e65ec698966c5b35052c72 dbname=d5e9m1htvn9vqg sslmode=require")

	if err := s.DB.Ping(); err != nil {
		log.Error(err.Error())
		return
	}
// my ngu
	fmt.Println("Connected to database....")
}

func (s *Sql) Close() {
	s.DB.Close()
}
