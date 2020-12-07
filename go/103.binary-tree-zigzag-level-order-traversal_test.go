package main

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			root: treeFromLevelOrder(1, 2, 3, 4, NA, NA, 5),
			want: [][]int{
				{1}, {3, 2}, {4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
