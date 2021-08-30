package main

import "testing"

func Test_longestArithSeqLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3, 4}, want: 4},
		{nums: []int{2, 2, 3, 4}, want: 3},
		{nums: []int{1, 1, 1, 2, 1, 1, 1}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestArithSeqLength(tt.nums); got != tt.want {
				t.Errorf("longestArithSeqLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
