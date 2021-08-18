package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rdurelli/loggingParser/src/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
