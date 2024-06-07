package usecase

import (
	"github.com/ropehapi/desafio-stone/internal/entity"
)

type GetPersonTreeUsecaseOutputDTO struct {
	Person entity.Person `json:"person"`
}

type GetPersonTreeUseCase struct {
	PersonRepository       entity.PersonRepositoryInterface
	RelationshipRepository entity.RelationshipRepositoryInterface
}

func NewGetPersonTreeUseCase(personRepository entity.PersonRepositoryInterface, relationshipRepository entity.RelationshipRepositoryInterface) *GetPersonTreeUseCase {
	return &GetPersonTreeUseCase{
		PersonRepository:       personRepository,
		RelationshipRepository: relationshipRepository,
	}
}

func (uc *GetPersonTreeUseCase) Execute(id string) (GetPersonTreeUsecaseOutputDTO, error) {
	person, _ := uc.PersonRepository.FindById(id)
	relationships, err := uc.buildTree(id, person)
	if err != nil {
		return GetPersonTreeUsecaseOutputDTO{}, err
	}
	person.Relationships = relationships

	return GetPersonTreeUsecaseOutputDTO{
		Person: *person,
	}, nil
}

func (uc *GetPersonTreeUseCase) buildTree(id string, person *entity.Person) ([]entity.Relationship, error) {
	person, _ = uc.PersonRepository.FindById(id)
	relationshipsIds, _ := uc.RelationshipRepository.GetRelationShipsIdsFromPersonId(id)

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
