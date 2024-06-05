package entity

type PersonRepositoryInterface interface {
	Save(p *Person) error
}
