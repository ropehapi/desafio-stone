//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/ropehapi/desafio-stone/internal/application/usecase"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"github.com/ropehapi/desafio-stone/internal/infra/database"
	"github.com/ropehapi/desafio-stone/internal/infra/web"
)

var setPersonRepositoryDependecy = wire.NewSet(
	database.NewPersonRepository,
	wire.Bind(new(entity.PersonRepositoryInterface), new(*database.PersonRepository)),
)

func NewCreatePersonUseCase(db *sql.DB) *usecase.CreatePersonUseCase {
	wire.Build(
		setPersonRepositoryDependecy,
		usecase.NewCreatePersonUseCase,
	)
	return &usecase.CreatePersonUseCase{}
}

func NewWebPersonHandler(db *sql.DB) *web.WebPersonHandler {
	wire.Build(
		setPersonRepositoryDependecy,
		web.NewWebPersonHandler,
	)
	return &web.WebPersonHandler{}
}
