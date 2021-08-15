package main

import (
	"reflect"
	"testing"
)

func Test_sufficientSubset(t *testing.T) {
	tests := []struct {
		name  string
		root  *TreeNode
		limit int
		want  *TreeNode
	}{
		{
			root:  treeFromLevelOrder(1, 2, -3, -5, NA, 4),
			limit: -1,
			want:  treeFromLevelOrder(1, NA, -3, 4),
		},

		{
			root:  treeFromLevelOrder(5, 4, 8, 11, NA, 17, 4, 7, 1, NA, NA, 5, 3),
			limit: 22,
			want:  treeFromLevelOrder(5, 4, 8, 11, NA, 17, 4, 7, NA, NA, NA, 5),
		},

		{
			root:  treeFromLevelOrder(1, 2, 3, 4, -99, -99, 7, 8, 9, -99, -99, 12, 13, -99, 14),
			limit: 1,
			want:  treeFromLevelOrder(1, 2, 3, 4, NA, NA, 7, 8, 9, NA, 14),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sufficientSubset(tt.root, tt.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sufficientSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
