package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rdurelli/loggingParser/src/bootstrap"
	"go.uber.org/fx"
)

//var (
//	db     *sql.DB
//	env    lib.Env
//	logger lib.Logger
//)
//
//func init() {
//	env = lib.NewEnv()
//	DB, err := sql.Open("mysql", env.DBUsername+":"+env.DBPassword+"@tcp("+env.DBHost+":"+env.DBPort+")/"+env.DBName+"?parseTime=true")
//	if err != nil {
//		log.Fatal("Something went wrong trying to connect to the database ", err)
//	}
//	err = DB.Ping()
//	if err != nil {
//		log.Fatal("Something went wrong trying to ping to the database ")
//	}
//	db = DB
//	logger = lib.NewLogger(env)
//}

func main() {
	//s := time.Now()
	//logs := make([]models.Log, 0)
	//readerLog := utils.NewReaderLog(env.LogPath, &logs, db, logger)
	//readerLog.Execute()
	////fmt.Println(logs)
	//fmt.Println(len(logs))
	//fmt.Println("\nTime taken - ", time.Since(s))
	fx.New(bootstrap.Module).Run()
}
