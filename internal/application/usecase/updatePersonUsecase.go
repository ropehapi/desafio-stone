package usecase

import "github.com/ropehapi/desafio-stone/internal/entity"

type UpdatePersonUsecase struct {
	personRepository entity.PersonRepositoryInterface
}

func NewUpdatePersonUsecase(personRepository entity.PersonRepositoryInterface) UpdatePersonUsecase {
	return UpdatePersonUsecase{
		personRepository: personRepository,
	}
}

func (uc UpdatePersonUsecase) Execute(id string, input CreatePersonUseCaseInputDTO) (PersonUseCaseOutputDTO, error) {
	person := &entity.Person{
		Name: input.Name,
	}

	if err := person.IsValid(); err != nil {
		return PersonUseCaseOutputDTO{}, err
	}

	if err := uc.personRepository.Update(id, person); err != nil {
		return PersonUseCaseOutputDTO{}, err
	}

	return PersonUseCaseOutputDTO{
		Id:   id,
		Name: person.Name,
	}, nil
}
