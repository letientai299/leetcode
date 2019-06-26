package main

import (
	"fmt"
	"testing"
)

func Test_isSubtree(t *testing.T) {
	tests := []struct {
		s    *TreeNode
		t    *TreeNode
		want bool
	}{
		{
			s:    treeFromLevelOrder(3, 4, 5, 1, 2),
			t:    treeFromLevelOrder(4, 1, 2),
			want: true,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.s, tt.t,
		)
		t.Run(testName, func(t *testing.T) {
			got := isSubtree(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isSubtree(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func Benchmark_isSubTree(b *testing.B) {
	s := treeFromLevelOrder(3, 4, 5, 1, 2)
	t := treeFromLevelOrder(4, 1, 2)
	for i := 0; i < b.N; i++ {
		_ = isSubtree(s, t)
	}
}
