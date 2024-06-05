package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ropehapi/desafio-stone/configs"
	webserver2 "github.com/ropehapi/desafio-stone/internal/infra/web/webserver"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", configs.DBUser, configs.DBPass, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	webserver := webserver2.NewWebServer(configs.WebServerPort)
	webPersonHandler := NewWebPersonHandler(db)
	webserver.RegisterRoutes("/person", webPersonHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Serve()
}
