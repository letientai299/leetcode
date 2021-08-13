package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSum(t *testing.T) {
	ms := Constructor677()
	ms.Insert("apple", 3)
	assert.Equal(t, ms.Sum("apple"), 3)
	assert.Equal(t, ms.Sum("appl"), 3)
	assert.Equal(t, ms.Sum("a"), 3)

	ms.Insert("app", 2)
	ms.Insert("apple", 2)
	assert.Equal(t, ms.Sum("ap"), 4)
}
