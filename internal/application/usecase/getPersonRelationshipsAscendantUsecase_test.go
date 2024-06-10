package usecase

import (
	"github.com/google/uuid"
	"github.com/ropehapi/desafio-stone/configs"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"github.com/ropehapi/desafio-stone/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPersonRelationshipsAscendantUsecase_Execute(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	personRepo := database.NewPersonRepository(tx)
	relRepo := database.NewRelationshipRepository(tx)

	filho := entity.NewPerson(uuid.New().String(), "Pedro Yoshimura")
	personRepo.Save(filho)

	pai := entity.NewPerson(uuid.New().String(), "Haruo Yoshimura")
	personRepo.Save(pai)

	paiDoPai := entity.NewPerson("aaa", "Shideharuo Yoshimura")
	personRepo.Save(paiDoPai)

	maeDoPai := entity.NewPerson("bbb", "Regina Querubin")
	personRepo.Save(maeDoPai)

	rel1 := entity.NewRelationship(filho, pai)
	relRepo.Save(rel1)

	rel2 := entity.NewRelationship(pai, paiDoPai)
	relRepo.Save(rel2)

	rel3 := entity.NewRelationship(pai, maeDoPai)
	relRepo.Save(rel3)

	getPersonTreeAscendantUsecase := NewGetPersonRelationshipsAscendantUsecase(personRepo, relRepo)
	output, err := getPersonTreeAscendantUsecase.Execute(filho.ID)
	assert.Nil(t, err)
	assert.Equal(t, output.Person.Relationships[0].Parent.Name, "Haruo Yoshimura")
	assert.Equal(t, output.Person.Relationships[0].Parent.Relationships[0].Parent.Name, "Shideharuo Yoshimura")
	assert.Equal(t, output.Person.Relationships[0].Parent.Relationships[1].Parent.Name, "Regina Querubin")
}
