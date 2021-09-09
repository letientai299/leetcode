package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	c := Constructor146(2)
	c.Put(1, 1)
	c.Put(2, 2)
	/*
		["LRUCache","put","put","get","get","put","get","get","get"]
		[[2],[2,1],[3,2],[3],[2],[4,3],[2],[3],[4]]
	*/
	assert.Equal(t, c.Get(1), 1)

	c.Put(3, 3)
	assert.Equal(t, c.Get(2), -1)

	//assert.Equal(t, c.head.val, 1)
	//assert.Equal(t, c.tail.val, 1)
}
