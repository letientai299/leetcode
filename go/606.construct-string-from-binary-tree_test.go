package main

import (
	"fmt"
	"testing"
)

func Test_tree2str(t *testing.T) {
	tests := []struct {
		t    *TreeNode
		want string
	}{
		{treeFromLevelOrder(1, 2, 3, 4), "1(2(4))(3)"},
		{treeFromLevelOrder(1, 2, 3, NA, 4), "1(2()(4))(3)"},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.t,
		)
		t.Run(testName, func(t *testing.T) {
			got := tree2str(tt.t)
			if got != tt.want {
				t.Errorf("tree2str(%v) = %v, want %v", tt.t, got, tt.want)
			}
		})
	}
}
