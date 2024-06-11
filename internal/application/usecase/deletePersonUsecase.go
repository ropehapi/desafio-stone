package usecase

import (
	"errors"
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
	_, err := uc.PersonRepository.FindById(id)
	if err != nil {
		return errors.New("person not found")
	}

	if err := uc.PersonRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
