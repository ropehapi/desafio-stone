package database

import (
	"database/sql"
	"github.com/ropehapi/desafio-stone/internal/entity"
)

type RelationshipRepository struct {
	tx *sql.Tx
}

func NewRelationshipRepository(tx *sql.Tx) *RelationshipRepository {
	return &RelationshipRepository{
		tx: tx,
	}
}

func (r *RelationshipRepository) GetParentIdsFromPersonId(id string) ([]string, error) {
	stmt, err := r.tx.Prepare("SELECT parent_id FROM relationship WHERE children_id=?")
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

func (r *RelationshipRepository) GetChildrenIdsFromPersonId(id string) ([]string, error) {
	stmt, err := r.tx.Prepare("SELECT children_id FROM relationship WHERE parent_id=?")
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

func (r *RelationshipRepository) Save(relationship *entity.Relationship) error {
	stmt, err := r.tx.Prepare("INSERT INTO relationship (children_id, parent_id) VALUES (?, ?)")
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

func (r *RelationshipRepository) Delete(childrenId, parentId string) error {
	stmt, err := r.tx.Prepare("DELETE FROM relationship WHERE children_id=? AND parent_id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(childrenId, parentId)
	if err != nil {
		return err
	}

	return nil
}
