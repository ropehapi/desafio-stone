package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func GetConn() *sql.DB {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Errorf("erro ao fazer ping no banco de dados: %s", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))))
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
