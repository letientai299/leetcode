package main

import "testing"

func Test_canArrange(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want bool
	}{
		{
			arr:  []int{8, 6, 3, 3},
			k:    5,
			want: false,
		},

		{
			arr:  []int{2, 2, 2, 2},
			k:    4,
			want: true,
		},

		{
			arr:  []int{-4, -7, 5, 2, 9, 1, 10, 4, -8, -3},
			k:    3,
			want: true,
		},

		{
			arr:  []int{-1, 1, -2, 2, -3, 3, -4, 4},
			k:    3,
			want: true,
		},

		{
			arr:  []int{3, 5, 1, 2, 3, 4},
			k:    10,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canArrange(tt.arr, tt.k); got != tt.want {
				t.Errorf("canArrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
