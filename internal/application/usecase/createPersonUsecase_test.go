package usecase

import (
	"github.com/ropehapi/desafio-stone/configs"
	"github.com/ropehapi/desafio-stone/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePersonUseCase_Execute(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	personRepo := database.NewPersonRepository(tx)

	input := CreateUpdatePersonUsecaseInputDTO{
		Name: "Pedro Yoshimura",
	}

	usecase := NewCreatePersonUseCase(personRepo)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.Id)
	assert.Equal(t, "Pedro Yoshimura", output.Name)
}
