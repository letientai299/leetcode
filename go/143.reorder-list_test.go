package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reorderList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{head: newList(), want: newList()},
		{head: newList(1), want: newList(1)},
		{head: newList(1, 2, 3, 4), want: newList(1, 4, 2, 3)},
		{head: newList(1, 2, 3, 4, 5), want: newList(1, 5, 2, 4, 3)},
		{head: newList(1, 2), want: newList(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.head)
			assert.Equal(t, tt.head, tt.want)
		})
	}
}
