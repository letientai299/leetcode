package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		sum  int
		want [][]int
	}{
		{
			root: treeFromLevelOrder(0, 1, 1),
			sum:  1,
			want: [][]int{{0, 1}, {0, 1}},
		},

		{
			root: treeFromLevelOrder(5),
			sum:  5,
			want: [][]int{{5}},
		},

		{
			root: treeFromLevelOrder(5, 4, 8, 11, NA, 13, 4, 7, 2, NA, NA, 5, 1),
			sum:  26,
			want: [][]int{{5, 8, 13}},
		},

		{
			root: treeFromLevelOrder(5, 4, 8, 11, NA, 13, 4, 7, 2, NA, NA, 5, 1),
			sum:  27,
			want: [][]int{{5, 4, 11, 7}},
		},

		{
			root: treeFromLevelOrder(5, 4, 8, 11, NA, 13, 4, 7, 2, NA, NA, 5, 1),
			sum:  220,
			want: nil,
		},

		{
			root: treeFromLevelOrder(5, 4, 8, 11, NA, 13, 4, 7, 2, NA, NA, 5, 1),
			sum:  22,
			want: [][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.root, tt.sum); !assert.Equal(t, got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
