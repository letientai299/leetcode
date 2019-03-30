package main

import (
	"fmt"
	"testing"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			got := sumOfLeftLeaves(tt.root)
			if got != tt.want {
				t.Errorf("sumOfLeftLeaves(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
