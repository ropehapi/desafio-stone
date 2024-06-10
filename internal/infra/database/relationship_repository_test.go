package database

import (
	"github.com/google/uuid"
	"github.com/ropehapi/desafio-stone/configs"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelationshipRepository_Save(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	personRepo := NewPersonRepository(tx)
	relRepo := NewRelationshipRepository(tx)

	children := entity.NewPerson(uuid.New().String(), "Pedro Yoshimura")
	personRepo.Save(children)

	parent := entity.NewPerson(uuid.New().String(), "Cesar Yoshimura")
	personRepo.Save(parent)

	rel := entity.NewRelationship(children, parent)
	err := relRepo.Save(rel)
	assert.Nil(t, err)
}

func TestRelationshipRepository_GetParentIdsFromPersonId(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	personRepo := NewPersonRepository(tx)
	relRepo := NewRelationshipRepository(tx)

	childrenUuid := uuid.New().String()
	children := entity.NewPerson(childrenUuid, "Pedro Yoshimura")
	personRepo.Save(children)

	parentUuid := uuid.New().String()
	parent := entity.NewPerson(parentUuid, "Cesar Yoshimura")
	personRepo.Save(parent)

	rel := entity.NewRelationship(children, parent)
	relRepo.Save(rel)

	parentIds, err := relRepo.GetParentIdsFromPersonId(childrenUuid)
	assert.Nil(t, err)
	assert.Equal(t, len(parentIds), 1)
	assert.Equal(t, parentIds[0], parentUuid)
}

func TestRelationshipRepository_GetChildrenIdsFromPersonId(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	personRepo := NewPersonRepository(tx)
	relRepo := NewRelationshipRepository(tx)

	childrenUuid := uuid.New().String()
	children := entity.NewPerson(childrenUuid, "Pedro Yoshimura")
	personRepo.Save(children)

	parentUuid := uuid.New().String()
	parent := entity.NewPerson(parentUuid, "Cesar Yoshimura")
	personRepo.Save(parent)

	rel := entity.NewRelationship(children, parent)
	relRepo.Save(rel)

	parentIds, err := relRepo.GetChildrenIdsFromPersonId(parentUuid)
	assert.Nil(t, err)
	assert.Equal(t, len(parentIds), 1)
	assert.Equal(t, parentIds[0], childrenUuid)
}

func TestRelationshipRepository_Delete(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	personRepo := NewPersonRepository(tx)
	relRepo := NewRelationshipRepository(tx)

	childrenUuid := uuid.New().String()
	children := entity.NewPerson(childrenUuid, "Pedro Yoshimura")
	personRepo.Save(children)

	parentUuid := uuid.New().String()
	parent := entity.NewPerson(parentUuid, "Cesar Yoshimura")
	personRepo.Save(parent)

	rel := entity.NewRelationship(children, parent)
	relRepo.Save(rel)

	err := relRepo.Delete(childrenUuid, parentUuid)
	assert.Nil(t, err)
}
