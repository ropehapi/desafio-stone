package entity

type PersonRepositoryInterface interface {
	Save(p *Person) error
	FindById(id string) (*Person, error)
	FindAll() ([]Person, error)
	Update(id string, p *Person) error
	Delete(id string) error
}

type RelationshipRepositoryInterface interface {
	GetParentIdsFromPersonId(id string) ([]string, error)
	GetChildrenIdsFromPersonId(id string) ([]string, error)
	Save(r *Relationship) error
	Delete(childrenId, parentId string) error
}
