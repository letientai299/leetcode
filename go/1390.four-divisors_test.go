package main

import "testing"

func Test_sumFourDivisors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{8, 2, 3, 4, 5, 6, 7, 1, 9, 10}, want: 45},
		{nums: []int{21, 4, 7}, want: 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumFourDivisors(tt.nums); got != tt.want {
				t.Errorf("sumFourDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
