package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamRoom(t *testing.T) {
	e := Constructor855(10)
	tests := []int{0, 9, 4, 2}

	for _, want := range tests {
		got := e.Seat()
		assert.Equal(t, want, got)
	}

	e.Leave(4)
	assert.Equal(t, 5, e.Seat())
	t.Log(e.stores)
	e.Leave(9)
	assert.Equal(t, 9, e.Seat())
	e.Leave(0)
	assert.Equal(t, 0, e.Seat())
	assert.Equal(t, 7, e.Seat())
}

func TestExamRoom_2(t *testing.T) {
	e := Constructor855(10)
	assert.Equal(t, 0, e.Seat())
	assert.Equal(t, 9, e.Seat())
	assert.Equal(t, 4, e.Seat())

	e.Leave(9)
	e.Leave(0)
	assert.Equal(t, 9, e.Seat())
}
