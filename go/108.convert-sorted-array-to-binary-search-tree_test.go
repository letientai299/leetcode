package main

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{},
		{[]int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1, 2, 3}},
		{[]int{1, 1, 2, 2, 2, 2, 2, 3}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := sortedArrayToBST(tt.nums)
			inOrder := got.inOrder()
			if !reflect.DeepEqual(inOrder, tt.nums) {
				t.Errorf("sortedArrayToBST() = %v, want %v (in order travesal)", got, tt.nums)
			}
		})
	}
}
