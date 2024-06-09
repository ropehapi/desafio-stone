package usecase

import "github.com/ropehapi/desafio-stone/internal/entity"

type GetPersonRelationshipsDescendantUsecase struct {
	PersonRepository       entity.PersonRepositoryInterface
	RelationshipRepository entity.RelationshipRepositoryInterface
}

func NewGetPersonRelationshipsDescendantUsecase(personRepository entity.PersonRepositoryInterface, relationshipRepository entity.RelationshipRepositoryInterface) *GetPersonRelationshipsDescendantUsecase {
	return &GetPersonRelationshipsDescendantUsecase{
		PersonRepository:       personRepository,
		RelationshipRepository: relationshipRepository,
	}
}

func (uc *GetPersonRelationshipsDescendantUsecase) Execute(id string) (*GetPersonTreeUsecaseOutputDTO, error) {
	person, _ := uc.PersonRepository.FindById(id)
	relationships, err := uc.buildTree(id, person)
	if err != nil {
		return nil, err
	}
	person.Relationships = relationships

	return &GetPersonTreeUsecaseOutputDTO{
		Person: *person,
	}, nil
}

func (uc *GetPersonRelationshipsDescendantUsecase) buildTree(id string, person *entity.Person) ([]entity.Relationship, error) {
	person, _ = uc.PersonRepository.FindById(id)
	relationshipsIds, _ := uc.RelationshipRepository.GetChildrenIdsFromPersonId(id)

	numberOfRelationships := len(relationshipsIds)
	relationships := make([]entity.Relationship, numberOfRelationships)

	if relationshipsIds != nil {
		for key, value := range relationshipsIds {
			children, _ := uc.PersonRepository.FindById(value)
			relationships[key] = entity.Relationship{Children: children}
			children.Relationships, _ = uc.buildTree(value, children)
		}
	} else {
		return nil, nil
	}
	return relationships, nil
}
