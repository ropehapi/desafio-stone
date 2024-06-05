package entity

import "errors"

type Relationship struct {
	Children *Person
	Parent   *Person
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
	for _, rel := range r.Parent.Relationships {
		visited := make(map[string]bool)
		if hasCycle(rel, visited) {
			return errors.New("cycle detected")
		}
	}
	return nil
}

// Função que verifica se existe um ciclo na relação
func hasCycle(rel Relationship, visited map[string]bool) bool {
	// Se o Parent for nulo, retornamos falso pois atingimos o final do relacionamento
	if rel.Parent == nil {
		return false
	}

	// Verifica se o ID do Parent já foi visitado
	if visited[rel.Parent.ID] {
		return true
	}

	// Marca o ID do Parent como visitado
	visited[rel.Parent.ID] = true

	// Cria uma nova relação para o próximo nó e chama recursivamente
	for _, nextRel := range rel.Parent.Relationships {
		if hasCycle(nextRel, visited) {
			return true
		}
	}

	// Após a verificação, removemos o ID do Parent do conjunto de visitados
	delete(visited, rel.Parent.ID)

	return false
}
