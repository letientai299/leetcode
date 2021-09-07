package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LockingTree(t *testing.T) {
	parent := []int{-1, 0, 0, 1, 1, 2, 2}
	lt := Constructor(parent)

	assert.Equal(t, true, lt.Lock(2, 2))
	assert.Equal(t, false, lt.Unlock(2, 3))
	assert.Equal(t, true, lt.Unlock(2, 2))
	assert.Equal(t, true, lt.Lock(4, 5))

	assert.Equal(t, true, lt.Upgrade(0, 1))
	assert.Equal(t, false, lt.Lock(0, 1))
}
