package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCalendar(t *testing.T) {
	ca := Constructor729()
	assert.Equal(t, true, ca.Book(10, 20))
	assert.Equal(t, false, ca.Book(15, 25))
	assert.Equal(t, true, ca.Book(20, 30))
}
