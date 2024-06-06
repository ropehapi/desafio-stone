package entity

type PersonRepositoryInterface interface {
	Save(p *Person) error
	FindById(id string) (*Person, error)
	FindAll() ([]Person, error)
	Update(id string, p *Person) error
}
