package usecase

import (
	"github.com/google/uuid"
	"github.com/ropehapi/desafio-stone/configs"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"github.com/ropehapi/desafio-stone/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRelationshipUseCase_Execute(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	personRepo := database.NewPersonRepository(tx)
	relRepo := database.NewRelationshipRepository(tx)

	children := entity.NewPerson(uuid.New().String(), "Pedro Yoshimura")
	personRepo.Save(children)

	parent := entity.NewPerson(uuid.New().String(), "Cesar Yoshimura")
	personRepo.Save(parent)

	input := CreateRelationshipInputDTO{
		ChildrenId: children.ID,
		ParentId:   parent.ID,
	}

	createRelationshipUsecase := NewCreateRelationshipUsecase(personRepo, relRepo)
	err := createRelationshipUsecase.Execute(input)
	assert.Nil(t, err)
}

func TestCreateRelationshipUseCase_Execute_has_cycle(t *testing.T) {
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

	input := CreateRelationshipInputDTO{
		ChildrenId: paiDoPai.ID,
		ParentId:   filho.ID,
	}

	createRelationshipUsecase := NewCreateRelationshipUsecase(personRepo, relRepo)
	err := createRelationshipUsecase.Execute(input)
	assert.NotNil(t, err)
	assert.Equal(t, "cycle detected", err.Error())
}

func TestCreateRelationshipUseCase_Execute_brothers(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	personRepo := database.NewPersonRepository(tx)
	relRepo := database.NewRelationshipRepository(tx)

	filho1 := entity.NewPerson(uuid.New().String(), "Pedro Yoshimura")
	personRepo.Save(filho1)

	filho2 := entity.NewPerson(uuid.New().String(), "Pietra Yoshimura")
	personRepo.Save(filho2)

	pai := entity.NewPerson(uuid.New().String(), "Haruo Yoshimura")
	personRepo.Save(pai)

	rel1 := entity.NewRelationship(filho1, pai)
	relRepo.Save(rel1)

	rel2 := entity.NewRelationship(filho2, pai)
	relRepo.Save(rel2)

	input := CreateRelationshipInputDTO{
		ChildrenId: filho1.ID,
		ParentId:   filho2.ID,
	}

	createRelationshipUsecase := NewCreateRelationshipUsecase(personRepo, relRepo)
	err := createRelationshipUsecase.Execute(input)
	assert.NotNil(t, err)
	assert.Equal(t, "brothers", err.Error())
}
