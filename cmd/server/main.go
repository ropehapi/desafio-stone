package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ropehapi/desafio-stone/configs"
	_ "github.com/ropehapi/desafio-stone/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

//@title API Árvore genealógica
//@version 1.0
//@desciption API desenvolvida para o case técnico do processo seletivo da stone

//@contact.name Pedro Yoshimura
//@contact.email ropehapi@gmail.com

// @host localhost:8000
// @BasePath /
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

	webPersonHandler := NewWebPersonHandler(db)
	webRelationshipHandler := NewWebRelationshipHandler(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Post("/person", webPersonHandler.Create)
	r.Get("/person/{id}", webPersonHandler.Get)
	r.Get("/person", webPersonHandler.List)
	r.Put("/person/{id}", webPersonHandler.Update)
	r.Delete("/person/{id}", webPersonHandler.Delete)

	r.Get("/relationship/{id}/asc", webRelationshipHandler.GetRelationshipsAscendant)
	r.Get("/relationship/{id}/desc", webRelationshipHandler.GetRelationshipsDescendant)
	r.Post("/relationship", webRelationshipHandler.Create)
	r.Delete("/relationship", webRelationshipHandler.Delete)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	err = http.ListenAndServe(configs.WebServerPort, r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting web server on port", configs.WebServerPort)
}
