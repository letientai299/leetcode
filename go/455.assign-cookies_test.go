package main

import (
	"fmt"
	"testing"
)

func Test_findContentChildren(t *testing.T) {
	tests := []struct {
		g    []int
		s    []int
		want int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2, 3}, []int{1, 2}, 2},
		{[]int{1, 2, 3}, []int{2, 2}, 2},
		{[]int{1, 2, 3}, []int{2, 3}, 2},
		{[]int{1, 2, 3}, []int{4, 2, 3}, 3},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 3},
		{[]int{1, 2, 4}, []int{1, 2, 3}, 2},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
		{[]int{1, 2, 2, 2, 3}, []int{1, 2, 2, 2}, 4},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.g, tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := findContentChildren(tt.g, tt.s)
			if got != tt.want {
				t.Errorf("findContentChildren(%v, %v) = %v, want %v", tt.g, tt.s, got, tt.want)
			}
		})
	}
}
