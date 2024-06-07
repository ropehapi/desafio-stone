package usecase

import "github.com/ropehapi/desafio-stone/internal/entity"

type GetPersonUseCase struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewGetPersonUseCase(personRepository entity.PersonRepositoryInterface) *GetPersonUseCase {
	return &GetPersonUseCase{
		PersonRepository: personRepository,
	}
}

func (uc *GetPersonUseCase) Execute(id string) (*PersonUseCaseOutputDTO, error) {
	person, err := uc.PersonRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return &PersonUseCaseOutputDTO{
		Id:   person.ID,
		Name: person.Name,
	}, nil
}
