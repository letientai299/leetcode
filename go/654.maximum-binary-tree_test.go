package main

import (
	"reflect"
	"testing"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{
			nums: []int{3, 2, 1, 6, 0, 5},
			want: treeFromLevelOrder(6, 3, 5, NA, 2, 0, NA, NA, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
