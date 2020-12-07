package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_flatten(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{root: treeFromLevelOrder(1, 2, 5, 3, 4, NA, 6), want: treeFromLevelOrder(1, NA, 2, NA, 3, NA, 4, NA, 5, NA, 6)},
		{root: treeFromLevelOrder(1, NA, 2, 3, 4), want: treeFromLevelOrder(1, NA, 2, NA, 3, NA, 4)},
		{root: treeFromLevelOrder(1), want: treeFromLevelOrder(1)},
		{root: treeFromLevelOrder(1, NA, 2), want: treeFromLevelOrder(1, NA, 2)},
		{root: treeFromLevelOrder(1, 2), want: treeFromLevelOrder(1, NA, 2)},
		{root: treeFromLevelOrder(1, 2, 3), want: treeFromLevelOrder(1, NA, 2, NA, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.root)
			if assert.Equal(t, tt.root, tt.want) {
				t.Logf("got=%v, want=%v", tt.root, tt.want)
			}
		})
	}
}
