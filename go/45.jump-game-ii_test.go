package main

import "testing"

func Test_jump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3}, want: 2},
		{nums: []int{2, 3, 0, 1, 4}, want: 2},
		{nums: []int{2, 3, 1, 1, 4}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
