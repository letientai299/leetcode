package main

import "testing"

func Test_sumNumbers(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{root: treeFromLevelOrder(1, 2, 3), want: 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
