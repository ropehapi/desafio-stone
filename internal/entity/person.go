package entity

import "errors"

type Person struct {
	ID            string
	Name          string
	Relationships []Relationship
}

func NewPerson(id, name string) *Person {
	return &Person{
		ID:   id,
		Name: name,
	}
}

func (p *Person) IsValid() error {
	if p.Name == "" {
		return errors.New("person name is required")
	}
	return nil
}
