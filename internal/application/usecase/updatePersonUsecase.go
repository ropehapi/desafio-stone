package usecase

import (
	"errors"
	"github.com/ropehapi/desafio-stone/internal/entity"
)

type UpdatePersonUsecase struct {
	personRepository entity.PersonRepositoryInterface
}

func NewUpdatePersonUsecase(personRepository entity.PersonRepositoryInterface) UpdatePersonUsecase {
	return UpdatePersonUsecase{
		personRepository: personRepository,
	}
}

func (uc UpdatePersonUsecase) Execute(id string, input CreateUpdatePersonUsecaseInputDTO) (*PersonUseCaseOutputDTO, error) {
	person := &entity.Person{
		Name: input.Name,
	}

	if err := person.IsValid(); err != nil {
		return nil, err
	}

	_, err := uc.personRepository.FindById(id)
	if err != nil {
		return nil, errors.New("person not found")
	}

	if err := uc.personRepository.Update(id, person); err != nil {
		return nil, err
	}

	return &PersonUseCaseOutputDTO{
		Id:   id,
		Name: person.Name,
	}, nil
}
