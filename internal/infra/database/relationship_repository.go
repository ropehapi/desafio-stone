package database

import (
	"database/sql"
	"github.com/ropehapi/desafio-stone/internal/entity"
)

type RelationshipRepository struct {
	DB *sql.DB
}

func NewRelationshipRepository(db *sql.DB) *RelationshipRepository {
	return &RelationshipRepository{
		DB: db,
	}
}

func (r *RelationshipRepository) Save(relationship *entity.Relationship) error {
	stmt, err := r.DB.Prepare("INSERT INTO relationship (children_id, parent_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(relationship.Children.ID, relationship.Parent.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RelationshipRepository) GetRelationShipsIdsFromPersonId(id string) ([]string, error) {
	stmt, err := r.DB.Prepare("SELECT parent_id FROM relationship WHERE children_id=?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	var relationshipsIds []string
	for rows.Next() {
		var relationshipId string
		err = rows.Scan(
			&relationshipId,
		)
		if err != nil {
			return nil, err
		}
		relationshipsIds = append(relationshipsIds, relationshipId)
	}

	return relationshipsIds, nil
}
