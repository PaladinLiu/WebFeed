package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)


func initDB() *sql.DB{
	DB, _ := sql.Open("mysql", "root:114514@tcp(127.0.0.1:3306)/web_feed_db?charset=utf8")
	if DB == nil {
		print("db connection failed.")
		return nil
	}

	DB.SetMaxIdleConns(20)
	DB.SetMaxOpenConns(20)

	//verify db connection
	err := DB.Ping()
	if err != nil{
		print("db connect failed.")
		return nil
	}
	return DB
}
