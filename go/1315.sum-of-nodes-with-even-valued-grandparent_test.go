package main

import "testing"

func Test_sumEvenGrandparent(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			root: treeFromLevelOrder(6, 7, 8, 2, 7, 1, 3, 9, NA, 1, 4, NA, NA, NA, 5),
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumEvenGrandparent(tt.root); got != tt.want {
				t.Errorf("sumEvenGrandparent() = %v, want %v", got, tt.want)
			}
		})
	}
}
