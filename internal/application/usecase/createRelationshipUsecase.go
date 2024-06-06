package usecase

import "github.com/ropehapi/desafio-stone/internal/entity"

type CreateRelationshipInputDTO struct {
	ChildrenId string `json:"childrenId"`
	ParentId   string `json:"parentId"`
}

type CreateRelationshipOutputDTO struct {
	ChildrenId string `json:"childrenId"`
	ParentId   string `json:"parentId"`
}

type CreateRelationshipUsecase struct {
	PersonRepository       entity.PersonRepositoryInterface
	RelationshipRepository entity.RelationshipRepositoryInterface
}

func NewCreateRelationshipUsecase(personRepository entity.PersonRepositoryInterface, relationshipRepository entity.RelationshipRepositoryInterface) CreateRelationshipUsecase {
	return CreateRelationshipUsecase{
		PersonRepository:       personRepository,
		RelationshipRepository: relationshipRepository,
	}
}

func (uc *CreateRelationshipUsecase) Execute(input CreateRelationshipInputDTO) (CreateRelationshipOutputDTO, error) {
	children, err := uc.PersonRepository.FindById(input.ChildrenId)
	if err != nil {
		return CreateRelationshipOutputDTO{}, err
	}

	parent, err := uc.PersonRepository.FindById(input.ParentId)
	if err != nil {
		return CreateRelationshipOutputDTO{}, err
	}

	relationship := entity.NewRelationship(children, parent)
	if err := relationship.IsValid(); err != nil {
		return CreateRelationshipOutputDTO{}, err
	}

	if err = uc.RelationshipRepository.Save(relationship); err != nil {
		return CreateRelationshipOutputDTO{}, err
	}

	return CreateRelationshipOutputDTO{
		ChildrenId: relationship.Children.ID,
		ParentId:   relationship.Parent.ID,
	}, err
}
