package lib

import (
	"database/sql"
	"log"
)

type Db struct {
	env Env
	Db  *sql.DB
}

func newDataBase(env Env) Db {
	DB, err := sql.Open("mysql", env.DBUsername+":"+env.DBPassword+"@tcp("+env.DBHost+":"+env.DBPort+")/"+env.DBName+"?parseTime=true")
	if err != nil {
		log.Fatal("Something went wrong trying to connect to the database ", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Something went wrong trying to ping to the database ")
	}
	return Db{
		Db: DB,
	}
}
