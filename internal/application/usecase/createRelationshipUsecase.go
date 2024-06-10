package usecase

import (
	"errors"
	"github.com/ropehapi/desafio-stone/internal/entity"
)

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

func (uc *CreateRelationshipUsecase) Execute(input CreateRelationshipInputDTO) (*CreateRelationshipOutputDTO, error) {
	if input.ChildrenId == input.ParentId {
		return nil, errors.New("Cycle detected")
	}

	children, err := uc.PersonRepository.FindById(input.ChildrenId)
	if err != nil {
		return nil, err
	}

	parent, err := uc.PersonRepository.FindById(input.ParentId)
	if err != nil {
		return nil, err
	}

	err = uc.validateCycle(children.ID, parent.ID, children)
	if err != nil {
		return nil, err
	}
	err = uc.validateBrothers(children.ID, parent.ID)
	if err != nil {
		return nil, err
	}

	relationship := entity.NewRelationship(children, parent)
	if err = relationship.IsValid(); err != nil {
		return nil, err
	}

	if err = uc.RelationshipRepository.Save(relationship); err != nil {
		return nil, err
	}

	return &CreateRelationshipOutputDTO{
		ChildrenId: relationship.Children.ID,
		ParentId:   relationship.Parent.ID,
	}, err
}

func (uc *CreateRelationshipUsecase) validateCycle(id, parentId string, person *entity.Person) error {
	person, _ = uc.PersonRepository.FindById(id)
	relationshipsIds, _ := uc.RelationshipRepository.GetChildrenIdsFromPersonId(id)

	if relationshipsIds != nil {
		for _, value := range relationshipsIds {
			if value == parentId {
				return errors.New("cycle detected")
			}
			children, _ := uc.PersonRepository.FindById(value)
			err := uc.validateCycle(value, parentId, children)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (uc *CreateRelationshipUsecase) validateBrothers(id, parentId string) error {
	parentIds, _ := uc.RelationshipRepository.GetParentIdsFromPersonId(id)
	if parentIds != nil {
		for _, value := range parentIds {
			childrenIds, _ := uc.RelationshipRepository.GetChildrenIdsFromPersonId(value)
			for _, childrenId := range childrenIds {
				if parentId == childrenId {
					return errors.New("brothers")
				}
			}
		}
	}

	return nil
}
