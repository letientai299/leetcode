package main

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		s    int
		nums []int
		want int
	}{
		{
			s:    3,
			nums: []int{1, 1},
			want: 0,
		},

		{
			s:    15,
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},

		{
			s:    7,
			nums: []int{2, 3, 1, 2, 4, 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minSubArrayLen(tt.s, tt.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
