package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Sql struct { 
	Db *sqlx.DB
	Host string
	Port int
	UserName string
	PassWord string
	DbName string
}

func (s *Sql) Connect(){
	dateSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	 s.Host , s.Port, s.UserName , s.PassWord, s.DbName)

	 s.Db = sqlx.MustConnect("postgres", dateSource)

	 if err := s.Db.Ping(); err != nil {
		log.Error(err.Error())
		return
	 }

	 fmt.Println("Connected Database")
}

func (s *Sql) Close(){
	s.Db.Close()
}