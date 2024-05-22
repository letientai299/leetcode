package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evaluateTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: treeFromLevelOrder(3, 3, 2, 0, 0, 3, 2, NA, NA, NA, NA, 1, 3, 1, 1, NA, NA, 2, 1, NA, NA, NA, NA, 1, 1, NA, NA, NA, NA, NA, NA),
			want: false,
		},

		{
			root: treeFromLevelOrder(2, 1, 3, NA, NA, 0, 1),
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equalf(t, tt.want, evaluateTree(tt.root), "evaluateTree(%v)", tt.root)
		})
	}
}
