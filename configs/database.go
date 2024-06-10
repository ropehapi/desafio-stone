package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetConn(configs Config) *sql.DB {
	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", configs.DBUser, configs.DBPass, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	return db
}

func GetTestConn() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/desafio_stone_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
