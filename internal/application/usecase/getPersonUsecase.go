package usecase

import "github.com/ropehapi/desafio-stone/internal/entity"

type GetPersonUseCaseInputDTO struct {
	ID string `json:"id"`
}

type GetPersonUseCase struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewGetPersonUseCase(personRepository entity.PersonRepositoryInterface) *GetPersonUseCase {
	return &GetPersonUseCase{
		PersonRepository: personRepository,
	}
}

func (uc *GetPersonUseCase) Execute(input GetPersonUseCaseInputDTO) (PersonUseCaseOutputDTO, error) {
	person, err := uc.PersonRepository.FindById(input.ID)
	if err != nil {
		return PersonUseCaseOutputDTO{}, err
	}

	return PersonUseCaseOutputDTO{
		Id:   person.ID,
		Name: person.Name,
	}, nil
}
