package main

import "testing"

func Test_arrayNesting(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 0},
			want: 3,
		},

		{
			nums: []int{1, 0},
			want: 2,
		},

		{
			nums: []int{0, 1},
			want: 1,
		},

		{
			nums: []int{0},
			want: 1,
		},

		{
			nums: []int{5, 4, 0, 3, 1, 6, 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayNesting(tt.nums); got != tt.want {
				t.Errorf("arrayNesting() = %v, want %v", got, tt.want)
			}
		})
	}
}
