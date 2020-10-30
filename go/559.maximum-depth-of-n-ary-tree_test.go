package main

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		root *Node
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
			got := maxDepth(tt.root)
			if got != tt.want {
				t.Errorf("maxDepth(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
