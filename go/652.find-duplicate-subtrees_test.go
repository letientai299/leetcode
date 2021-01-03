package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDuplicateSubtrees(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []*TreeNode
	}{
		{
			root: treeFromLevelOrder(0, 0, 0, 0, NA, NA, 0, 0, 0, 0, 0),
			want: []*TreeNode{
				treeFromLevelOrder(0, 0, 0),
				treeFromLevelOrder(0),
			},
		},

		{
			root: treeFromLevelOrder(1, 2, 3, 4, NA, 2, 4, NA, NA, 4),
			want: []*TreeNode{
				treeFromLevelOrder(2, 4),
				treeFromLevelOrder(4),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicateSubtrees(tt.root); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("findDuplicateSubtrees() = %v, want %v", got, tt.want)
				tt.root.print()
			}
		})
	}
}
