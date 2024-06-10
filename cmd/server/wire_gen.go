// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/ropehapi/desafio-stone/internal/application/usecase"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"github.com/ropehapi/desafio-stone/internal/infra/database"
	"github.com/ropehapi/desafio-stone/internal/infra/web"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreatePersonUseCase(tx *sql.Tx) *usecase.CreatePersonUseCase {
	personRepository := database.NewPersonRepository(tx)
	createPersonUseCase := usecase.NewCreatePersonUseCase(personRepository)
	return createPersonUseCase
}

func NewWebPersonHandler(tx *sql.Tx) *web.WebPersonHandler {
	personRepository := database.NewPersonRepository(tx)
	webPersonHandler := web.NewWebPersonHandler(personRepository)
	return webPersonHandler
}

func NewWebRelationshipHandler(tx *sql.Tx) *web.WebRelationshipHandler {
	personRepository := database.NewPersonRepository(tx)
	relationshipRepository := database.NewRelationshipRepository(tx)
	webRelationshipHandler := web.NewWebRelationshipHandler(personRepository, relationshipRepository)
	return webRelationshipHandler
}

// wire.go:

var setPersonRepositoryDependecy = wire.NewSet(database.NewPersonRepository, wire.Bind(new(entity.PersonRepositoryInterface), new(*database.PersonRepository)))

var setRelationshipRepositoryDependecy = wire.NewSet(database.NewPersonRepository, database.NewRelationshipRepository, wire.Bind(new(entity.PersonRepositoryInterface), new(*database.PersonRepository)), wire.Bind(new(entity.RelationshipRepositoryInterface), new(*database.RelationshipRepository)))
