package main

import "testing"

func Test_isSubPath(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		root *TreeNode
		want bool
	}{
		{
			root: treeFromLevelOrder(1, 4, 4, NA, 2, 2, NA, 1, NA, 6, 8, NA, NA, NA, NA, 1, 3),
			head: newList(4, 2, 1),
			want: true,
		},

		{
			root: treeFromLevelOrder(1, 4, 4, NA, 2, 2, NA, 1, NA, 6, 8, NA, NA, NA, NA, 1, 3),
			head: newList(4, 2, 7),
			want: false,
		},

		{
			root: treeFromLevelOrder(1, 4, 4, NA, 2, 2, NA, 1, NA, 6, 8, NA, NA, NA, NA, 1, 3),
			head: newList(4, 2, 8),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubPath(tt.head, tt.root); got != tt.want {
				t.Errorf("isSubPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
