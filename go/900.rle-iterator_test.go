package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRLEIterator_Next(t *testing.T) {
	it := Constructor([]int{3, 8, 0, 9, 2, 5})
	assert.Equal(t, 8, it.Next(2))
	assert.Equal(t, 8, it.Next(1))
	assert.Equal(t, 5, it.Next(1))
	assert.Equal(t, -1, it.Next(2))
}
