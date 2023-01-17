package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "mock:password@tcp(127.0.0.1:4407)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
