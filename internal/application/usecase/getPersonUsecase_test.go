package usecase

import (
	"github.com/ropehapi/desafio-stone/configs"
	"github.com/ropehapi/desafio-stone/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPersonUseCase_Execute(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	personRepo := database.NewPersonRepository(tx)

	input := CreateUpdatePersonUsecaseInputDTO{
		Name: "Pedro Yoshimura",
	}

	createPersonUsecase := NewCreatePersonUseCase(personRepo)
	output, err := createPersonUsecase.Execute(input)

	getPersonUsecase := NewGetPersonUseCase(personRepo)
	output, err = getPersonUsecase.Execute(output.Id)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, output.Name, input.Name)
}
