package entity

type PersonRepositoryInterface interface {
	Save(p *Person) error
	FindById(id string) (*Person, error)
	FindAll() ([]Person, error)
	Update(id string, p *Person) error
	Delete(id string) error
	GetRelationShipsIds(id string) ([]string, error)
}

type RelationshipRepositoryInterface interface {
	Save(r *Relationship) error
}
