package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularQueue(t *testing.T) {
	k := 3
	q := Constructor(k)
	assert.True(t, q.IsEmpty())

	for i := 0; i < k; i++ {
		v := i*10 + 1
		assert.Truef(t, q.EnQueue(v), "enqueue %d should ok", v)
		assert.Equal(t, v, q.Rear())
	}

	assert.True(t, q.IsFull())
	assert.False(t, q.EnQueue(100*10))

	for i := 0; i < k; i++ {
		v := i*10 + 1
		assert.Equal(t, v, q.Front())
		assert.True(t, q.DeQueue())
	}

	assert.False(t, q.IsFull())
	assert.True(t, q.IsEmpty())
}
