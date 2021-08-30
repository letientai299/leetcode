package main

import "testing"

func Test_btreeGameWinningMove(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		n    int
		x    int
		want bool
	}{
		{
			root: treeFromLevelOrder(1, 2, 3, 4, 5),
			n:    5,
			x:    1,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := btreeGameWinningMove(tt.root, tt.n, tt.x); got != tt.want {
				t.Errorf("btreeGameWinningMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
