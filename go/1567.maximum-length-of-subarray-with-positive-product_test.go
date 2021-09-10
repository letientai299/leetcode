package main

import "testing"

func Test_getMaxLen(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{0, 1, -2, -3, -4}, want: 3},
		{nums: []int{1, -2, -3, 4}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxLen(tt.nums); got != tt.want {
				t.Errorf("getMaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
