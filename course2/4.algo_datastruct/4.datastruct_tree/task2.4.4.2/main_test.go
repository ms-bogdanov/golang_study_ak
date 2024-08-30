package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	bt := NewBTree(2)
	user := User{ID: 1, Name: "Alice", Age: 30}
	bt.Insert(user)

	result := bt.Search(1)

	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Age, result.Age)

	result = bt.Search(-1)

	assert.Nil(t, result)
}
