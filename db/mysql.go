package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	MySQL *sql.DB
	err   error
)

//it will init on server start up
func init() {
	MySQL, err = sql.Open("mysql", "root:admin@/dev?charset=utf8")

	if err != nil {
		log.Fatalf("mysql open error: %v", err.Error())
	}

	if err = MySQL.Ping(); err != nil {
		log.Fatalf("mysql connection error: %v", err.Error())
	}

	MySQL.SetMaxIdleConns(10)
	MySQL.SetMaxOpenConns(10)
}
