package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	tests := []struct {
		node *ListNode
		want *ListNode
	}{
		{
			node: newList(1, 2, 3, 4),
			want: newList(2, 3, 4),
		},
		{
			node: newList(1, 2),
			want: newList(2),
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.node,
		)
		t.Run(testName, func(t *testing.T) {
			deleteNode(tt.node)
			assert.Equal(t, tt.node, tt.want)
		})
	}
}
