package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_upsideDownBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			root: treeFromLevelOrder(1, 2),
			want: treeFromLevelOrder(2, NA, 1),
		},

		{
			root: treeFromLevelOrder(1, 2, 3),
			want: treeFromLevelOrder(2, 3, 1),
		},

		{
			root: treeFromLevelOrder(1, 2, 3, 4, 5),
			want: treeFromLevelOrder(4, 5, 2, NA, NA, 3, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, upsideDownBinaryTree(tt.root), "upsideDownBinaryTree(%v)", tt.root)
		})
	}
}
