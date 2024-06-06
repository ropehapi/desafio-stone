package usecase

import (
	"github.com/google/uuid"
	"github.com/ropehapi/desafio-stone/internal/entity"
)

type CreatePersonUseCaseInputDTO struct {
	Name string `json:"name"`
}

type PersonUseCaseOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreatePersonUseCase struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewCreatePersonUseCase(personRepository entity.PersonRepositoryInterface) *CreatePersonUseCase {
	return &CreatePersonUseCase{
		PersonRepository: personRepository,
	}
}

func (uc *CreatePersonUseCase) Execute(input CreatePersonUseCaseInputDTO) (PersonUseCaseOutputDTO, error) {
	person := &entity.Person{
		ID:   uuid.New().String(),
		Name: input.Name,
	}
	if err := person.IsValid(); err != nil {
		return PersonUseCaseOutputDTO{}, err
	}

	if err := uc.PersonRepository.Save(person); err != nil {
		return PersonUseCaseOutputDTO{}, err
	}

	return PersonUseCaseOutputDTO{
		Id:   person.ID,
		Name: person.Name,
	}, nil
}
