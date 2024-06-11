package usecase

import (
	"errors"
	"github.com/ropehapi/desafio-stone/internal/entity"
)

type GetPersonTreeUsecaseOutputDTO struct {
	Person entity.Person `json:"person"`
}

type GetPersonRelationshipsAscendantUsecase struct {
	PersonRepository       entity.PersonRepositoryInterface
	RelationshipRepository entity.RelationshipRepositoryInterface
}

func NewGetPersonRelationshipsAscendantUsecase(personRepository entity.PersonRepositoryInterface, relationshipRepository entity.RelationshipRepositoryInterface) *GetPersonRelationshipsAscendantUsecase {
	return &GetPersonRelationshipsAscendantUsecase{
		PersonRepository:       personRepository,
		RelationshipRepository: relationshipRepository,
	}
}

func (uc *GetPersonRelationshipsAscendantUsecase) Execute(id string) (*GetPersonTreeUsecaseOutputDTO, error) {
	person, _ := uc.PersonRepository.FindById(id)
	if person == nil {
		return nil, errors.New("person not found")
	}

	relationships, err := uc.buildTree(id, person)
	if err != nil {
		return nil, err
	}
	person.Relationships = relationships

	return &GetPersonTreeUsecaseOutputDTO{
		Person: *person,
	}, nil
}

func (uc *GetPersonRelationshipsAscendantUsecase) buildTree(id string, person *entity.Person) ([]entity.Relationship, error) {
	person, _ = uc.PersonRepository.FindById(id)
	relationshipsIds, _ := uc.RelationshipRepository.GetParentIdsFromPersonId(id)

	numberOfRelationships := len(relationshipsIds)
	relationships := make([]entity.Relationship, numberOfRelationships)

	if relationshipsIds != nil {
		for key, value := range relationshipsIds {
			parent, _ := uc.PersonRepository.FindById(value)
			relationships[key] = entity.Relationship{Parent: parent}
			parent.Relationships, _ = uc.buildTree(value, parent)
		}
	} else {
		return nil, nil
	}
	return relationships, nil
}
