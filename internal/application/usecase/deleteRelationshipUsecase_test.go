package usecase

import (
	"github.com/google/uuid"
	"github.com/ropehapi/desafio-stone/configs"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"github.com/ropehapi/desafio-stone/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteRelationshipUseCase_Execute(t *testing.T) {
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

	rel := entity.NewRelationship(children, parent)
	relRepo.Save(rel)

	input := DeleteRelationshipUsecaseInputDTO{
		ChildrenId: children.ID,
		ParentId:   parent.ID,
	}

	deleteRelationshipUsecase := NewDeleteRelationshipUsecase(relRepo)
	err := deleteRelationshipUsecase.Execute(input)
	assert.Nil(t, err)
}
