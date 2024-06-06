package usecase

import "github.com/ropehapi/desafio-stone/internal/entity"

type ListPersonUseCase struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewListPersonUseCase(personRepository entity.PersonRepositoryInterface) ListPersonUseCase {
	return ListPersonUseCase{
		PersonRepository: personRepository,
	}
}

func (uc ListPersonUseCase) Execute() ([]PersonUseCaseOutputDTO, error) {
	var list []PersonUseCaseOutputDTO

	persons, err := uc.PersonRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, person := range persons {
		list = append(list, PersonUseCaseOutputDTO{
			Id:   person.ID,
			Name: person.Name,
		})
	}

	return list, nil
}
