package main

import "testing"

func Test_distributeCoins(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{root: treeFromLevelOrder(1, 0, 0, NA, 3), want: 4},
		{root: treeFromLevelOrder(3, 0, 0), want: 2},
		{root: treeFromLevelOrder(1, 1, 1), want: 0},
		{root: treeFromLevelOrder(1, 0, 2), want: 2},
		{root: treeFromLevelOrder(0, 0, 3), want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCoins(tt.root); got != tt.want {
				t.Errorf("distributeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
