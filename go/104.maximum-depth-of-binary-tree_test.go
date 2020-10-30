package main

import (
	"fmt"
	"testing"
)

func Test_maxDepth_104(t *testing.T) {
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
			got := maxDepth_104(tt.root)
			if got != tt.want {
				t.Errorf("maxDepth(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
