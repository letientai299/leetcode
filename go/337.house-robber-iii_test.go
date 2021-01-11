package main

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			/*
						 3
				  2    3
				 * 3  * 1
			*/
			root: treeFromLevelOrder(3, 2, 3, NA, 3, NA, 1),
			want: 7,
		},

		{
			root: treeFromLevelOrder(4, 1, NA, 2, NA, 3),
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.root); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
