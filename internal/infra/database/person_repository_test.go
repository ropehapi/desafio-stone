package database

import (
	"github.com/google/uuid"
	"github.com/ropehapi/desafio-stone/configs"
	"github.com/ropehapi/desafio-stone/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonRepositorySaveAndFindById(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	repo := NewPersonRepository(tx)

	uuid := uuid.New().String()
	person := entity.NewPerson(uuid, "Pedro Yoshimura")
	err := repo.Save(person)
	assert.Nil(t, err)

	personDatabase, err := repo.FindById(uuid)
	assert.NotNil(t, personDatabase)
	assert.Equal(t, "Pedro Yoshimura", personDatabase.Name)
}

func TestPersonRepository_FindAll(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	repo := NewPersonRepository(tx)

	persons := []entity.Person{
		{
			ID:   uuid.New().String(),
			Name: "Pedro Yoshimura",
		},
		{
			ID:   uuid.New().String(),
			Name: "Pedro Oshimura",
		},
		{
			ID:   uuid.New().String(),
			Name: "Pedro Ioshimura",
		},
	}

	for _, person := range persons {
		repo.Save(&person)
	}

	personsResult, err := repo.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, len(persons), len(personsResult))
}

func TestPersonRepositoryUpdate(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	repo := NewPersonRepository(tx)

	uuid := uuid.New().String()
	person := entity.NewPerson(uuid, "Pedro Yoshimura")

	repo.Save(person)

	person.Name = "Pedro Oshimura"
	err := repo.Update(uuid, person)
	assert.Nil(t, err)

	personDatabase, err := repo.FindById(uuid)
	assert.Nil(t, err)
	assert.Equal(t, "Pedro Oshimura", personDatabase.Name)
}

func TestPersonRepositoryDelete(t *testing.T) {
	db := configs.GetTestConn()
	defer db.Close()

	tx, _ := db.Begin()
	defer tx.Rollback()

	repo := NewPersonRepository(tx)

	uuid := uuid.New().String()
	person := entity.NewPerson(uuid, "Pedro Yoshimura")
	repo.Save(person)

	err := repo.Delete(uuid)
	assert.Nil(t, err)

	personDatabase, err := repo.FindById(uuid)
	assert.NotNil(t, err)
	assert.Nil(t, personDatabase)
}
