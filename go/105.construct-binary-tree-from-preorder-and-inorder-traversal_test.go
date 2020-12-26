package main

import (
	"reflect"
	"testing"
)

func Test_buildTree105(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			preorder: []int{1, 2, 3, 4, 5},
			inorder:  []int{5, 4, 3, 2, 1},
			want:     treeFromLevelOrder(1, 2, NA, 3, NA, 4, NA, 5),
		},

		{
			preorder: []int{},
			inorder:  []int{},
			want:     nil,
		},

		{
			preorder: []int{1, 3, 5, 6},
			inorder:  []int{1, 5, 3, 6},
			want:     treeFromLevelOrder(1, NA, 3, 5, 6),
		},

		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want:     treeFromLevelOrder(3, 9, 20, NA, NA, 15, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree105(tt.preorder, tt.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
