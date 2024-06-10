package usecase

import (
	"github.com/ropehapi/desafio-stone/configs"
	"github.com/ropehapi/desafio-stone/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePersonUseCase_Update(t *testing.T) {
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

	input = CreateUpdatePersonUsecaseInputDTO{
		Name: "Pedro Oshimura",
	}
	deletePersonUsecase := NewUpdatePersonUsecase(personRepo)
	output, err = deletePersonUsecase.Execute(output.Id, input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, output.Name, input.Name)
}
