package main

import "testing"

func Test_check(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{nums: []int{3, 4, 5, 1, 2}, want: true},
		{nums: []int{3, 6, 5, 1, 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.nums); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
