package database

import (
	"database/sql"
	"github.com/ropehapi/desafio-stone/internal/entity"
)

type PersonRepository struct {
	DB *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{
		DB: db,
	}
}

func (r *PersonRepository) Save(person *entity.Person) error {
	stmt, err := r.DB.Prepare("INSERT INTO person (id, name) values(?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(person.ID, person.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) FindById(id string) (*entity.Person, error) {
	stmt, err := r.DB.Prepare("SELECT id, name FROM person WHERE id=?")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)
	var person entity.Person
	err = row.Scan(
		&person.ID,
		&person.Name,
	)
	if err != nil {
		return nil, err
	}

	return &person, nil
}
