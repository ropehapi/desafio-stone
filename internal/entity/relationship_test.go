package entity

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRelationship(t *testing.T) {
	child := NewPerson("1", "Child")
	parent := NewPerson("2", "Parent")
	relationship := NewRelationship(child, parent)

	assert.NotNil(t, relationship, "NewRelationship should return a valid Relationship instance")
	assert.Equal(t, child, relationship.Children, "Relationship Children should be set correctly")
	assert.Equal(t, parent, relationship.Parent, "Relationship Parent should be set correctly")
}

func TestRelationship_IsValid(t *testing.T) {
	child := NewPerson("1", "Child")
	parent := NewPerson("2", "Parent")

	// Test case: valid relationship
	relationship := NewRelationship(child, parent)
	err := relationship.IsValid()
	assert.NoError(t, err, "Valid relationship should not return an error")

	// Test case: relationship with nil children
	relationship = NewRelationship(nil, parent)
	err = relationship.IsValid()
	require.Error(t, err, "Relationship with nil children should return an error")
	assert.Equal(t, "relationship children is empty", err.Error(), "Error message should be 'relationship children is empty'")

	// Test case: relationship with nil parent
	relationship = NewRelationship(child, nil)
	err = relationship.IsValid()
	require.Error(t, err, "Relationship with nil parent should return an error")
	assert.Equal(t, "relationship parent is empty", err.Error(), "Error message should be 'relationship parent is empty'")
}
