package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findMode(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{treeFromLevelOrder(1, NA, 2, 2), []int{2}},
		{treeFromLevelOrder(1, 1, 2, 1, 2), []int{1}},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.root,
		)
		t.Run(testName, func(t *testing.T) {
			if got := findMode(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
