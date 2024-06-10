package database

import (
	"database/sql"
	"github.com/ropehapi/desafio-stone/internal/entity"
)

type PersonRepository struct {
	tx *sql.Tx
}

func NewPersonRepository(tx *sql.Tx) *PersonRepository {
	return &PersonRepository{
		tx: tx,
	}
}

func (r *PersonRepository) Save(person *entity.Person) error {
	stmt, err := r.tx.Prepare("INSERT INTO person (id, name) values(?,?)")
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
	stmt, err := r.tx.Prepare("SELECT id, name FROM person WHERE id=?")
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

func (r *PersonRepository) FindAll() ([]entity.Person, error) {
	stmt, err := r.tx.Prepare("SELECT id, name FROM person")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var persons []entity.Person
	for rows.Next() {
		var person entity.Person
		err = rows.Scan(
			&person.ID,
			&person.Name,
		)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	return persons, nil
}

func (r *PersonRepository) Update(id string, p *entity.Person) error {
	stmt, err := r.tx.Prepare("UPDATE person SET name=? WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) Delete(id string) error {
	stmt, err := r.tx.Prepare("DELETE FROM person WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
