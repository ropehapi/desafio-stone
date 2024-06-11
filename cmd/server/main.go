package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/ropehapi/desafio-stone/configs"
	_ "github.com/ropehapi/desafio-stone/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"time"
)

//@title API Árvore genealógica
//@version 1.0
//@desciption API desenvolvida para o case técnico do processo seletivo da stone

//@contact.name Pedro Yoshimura
//@contact.email ropehapi@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := configs.GetConn()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Commit()

	webPersonHandler := NewWebPersonHandler(tx)
	webRelationshipHandler := NewWebRelationshipHandler(tx)

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

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	err = http.ListenAndServe(os.Getenv("WEB_SERVER_PORT"), r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting web server on port", os.Getenv("WEB_SERVER_PORT"))
}
