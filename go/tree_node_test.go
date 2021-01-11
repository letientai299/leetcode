package main

import (
	"reflect"
	"testing"
)

func Test_treeFromLevelOrder(t *testing.T) {
	tests := []struct {
		values []int
	}{
		{[]int{10, 5}},
		{[]int{10, 5, 3}},
		{[]int{10, 5, 3, 1}},
		{[]int{10, 5, -3, 3, 2, NA, 11, 3, -2, NA, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tree := treeFromLevelOrder(tt.values...)
			levelOderAgain := tree.levelOrder()

			if !reflect.DeepEqual(levelOderAgain, tt.values) {
				t.Errorf("treeFromLevelOrder() = %v, want %v", levelOderAgain, tt.values)
			}
		})
	}
}
