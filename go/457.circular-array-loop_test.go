package main

import "testing"

func Test_circularArrayLoop(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{nums: []int{1, 1}, want: true},
		{nums: []int{-2, -3, -9}, want: false},
		{nums: []int{2, -1, 2, -1, 3}, want: true},
		{nums: []int{3, 1, 2}, want: true},
		{nums: []int{2, 2, 2, 2, 2, 4, 7}, want: false},
		{nums: []int{-1, -2, -3, -4, -5}, want: false},
		{nums: []int{-1, 2}, want: false},
		{nums: []int{-2, 1, -1, -2, -2}, want: false},
		{nums: []int{2, -1, 1, 2, 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayLoop(tt.nums); got != tt.want {
				t.Errorf("circularArrayLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
