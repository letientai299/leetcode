package main

import "testing"

func Test_threeSumMulti(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
			target: 9,
			want:   20,
		},

		{
			nums:   []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
			target: 8,
			want:   20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumMulti(tt.nums, tt.target); got != tt.want {
				t.Errorf("threeSumMulti() = %v, want %v", got, tt.want)
			}
		})
	}
}
