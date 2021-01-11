package main

import "testing"

func Test_closestValue(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		target float64
		want   int
	}{
		{
			root:   treeFromLevelOrder(5, 3, 6, 2, 4, NA, NA, 1),
			target: 2.85,
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := closestValue(tt.root, tt.target); got != tt.want {
				t.Errorf("closestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
