package usecase

import "github.com/ropehapi/desafio-stone/internal/entity"

type DeleteRelationshipUsecaseInputDTO struct {
	ChildrenId string `json:"childrenId"`
	ParentId   string `json:"parentId"`
}

type DeleteRelationshipUsecase struct {
	RelationshipRepository entity.RelationshipRepositoryInterface
}

func NewDeleteRelationshipUsecase(deleteRelationshipRepository entity.RelationshipRepositoryInterface) *DeleteRelationshipUsecase {
	return &DeleteRelationshipUsecase{
		RelationshipRepository: deleteRelationshipRepository,
	}
}

func (uc *DeleteRelationshipUsecase) Execute(input DeleteRelationshipUsecaseInputDTO) error {
	if err := uc.RelationshipRepository.Delete(input.ChildrenId, input.ParentId); err != nil {
		return err
	}

	return nil
}
