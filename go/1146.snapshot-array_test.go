package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshotArray(t *testing.T) {
	sa := Constructor(1)
	assert.Equal(t, 0, sa.Snap())
	assert.Equal(t, 1, sa.Snap())
	sa.Set(0, 4)
	assert.Equal(t, 2, sa.Snap())

	assert.Equal(t, 0, sa.Get(0, 1))
	sa.Set(0, 12)
	assert.Equal(t, 0, sa.Get(0, 1))
	assert.Equal(t, 3, sa.Snap())

	assert.Equal(t, 12, sa.Get(0, 3))
}
