package main

import "testing"

func Test_wiggleMaxLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{3}, want: 1},
		{nums: []int{3}, want: 1},
		{nums: []int{3, 3, 3, 2, 5}, want: 3},
		{nums: []int{1, 2, 3, 4, 5}, want: 2},
		{nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wiggleMaxLength(tt.nums); got != tt.want {
				t.Errorf("wiggleMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
