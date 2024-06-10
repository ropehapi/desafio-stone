package usecase

import (
	"github.com/ropehapi/desafio-stone/configs"
	"github.com/ropehapi/desafio-stone/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePersonUseCase_List(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	personRepo := database.NewPersonRepository(tx)

	inputs := []CreateUpdatePersonUsecaseInputDTO{
		{
			Name: "Pedro Yoshimura",
		},
		{
			Name: "Pedro Oshimura",
		},
		{
			Name: "Pedro Ioshimura",
		},
	}

	createPersonUsecase := NewCreatePersonUseCase(personRepo)
	for _, input := range inputs {
		createPersonUsecase.Execute(input)
	}

	listPersonUsecase := NewListPersonUseCase(personRepo)
	output, err := listPersonUsecase.Execute()
	assert.Nil(t, err)
	assert.NotEmptyf(t, output, "output is empty")
	assert.NotNil(t, output)
}
