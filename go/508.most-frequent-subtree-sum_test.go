package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findFrequentTreeSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			root: treeFromLevelOrder(5, 2, -5),
			want: []int{2},
		},

		{
			root: treeFromLevelOrder(5, 2, -3),
			want: []int{4, 2, -3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFrequentTreeSum(tt.root); !assert.ElementsMatch(t, tt.want, got) {
				t.Errorf("findFrequentTreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
