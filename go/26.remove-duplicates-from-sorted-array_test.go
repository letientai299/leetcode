package main

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1, 1, 1, 1, 1, 1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5, 6}, 6},
		{[]int{-1, 2, 3, 4, 5, 6}, 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			fmt.Println(tt.nums[:got])
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
