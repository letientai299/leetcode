package main

import "testing"

func Test_triangleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{3, 4, 5, 6}, want: 4},
		{nums: []int{1, 2, 2, 2}, want: 4},
		{nums: []int{2, 2, 3, 4}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleNumber(tt.nums); got != tt.want {
				t.Errorf("triangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
