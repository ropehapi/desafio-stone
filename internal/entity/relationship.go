package entity

import "errors"

type Relationship struct {
	Children *Person `json:"children,omitempty"`
	Parent   *Person `json:"parent,omitempty"`
}

func NewRelationship(children, parent *Person) *Relationship {
	return &Relationship{
		Children: children,
		Parent:   parent,
	}
}

func (r *Relationship) IsValid() error {
	if r.Children == nil {
		return errors.New("relationship children is empty")
	}
	if r.Parent == nil {
		return errors.New("relationship parent is empty")
	}
	return nil
}
