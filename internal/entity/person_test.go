package entity

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPerson(t *testing.T) {
	id := "aaabbbccc"
	name := "Pedro Yoshimura"
	person := NewPerson(id, name)

	assert.NotNil(t, person, "NewPerson should return a valid Person instance")
	assert.Equal(t, id, person.ID, "Person ID should be set correctly")
	assert.Equal(t, name, person.Name, "Person Name should be set correctly")
	assert.Empty(t, person.Relationships, "New Person should have no relationships by default")
}

func TestPerson_IsValid(t *testing.T) {
	person := &Person{Name: "Pedro Yoshimura"}
	err := person.IsValid()
	assert.NoError(t, err, "Valid person should not return an error")

	person = &Person{Name: ""}
	err = person.IsValid()
	require.Error(t, err, "Invalid person should return an error")
	assert.Equal(t, "person name is required", err.Error(), "Error message should be 'person name is required'")
}
