package usecase

import (
	"github.com/ropehapi/desafio-stone/internal/entity"
)

type DeletePersonUseCase struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewDeletePersonUseCase(personRepository entity.PersonRepositoryInterface) DeletePersonUseCase {
	return DeletePersonUseCase{
		PersonRepository: personRepository,
	}
}

func (uc DeletePersonUseCase) Execute(id string) error {
	if err := uc.PersonRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
