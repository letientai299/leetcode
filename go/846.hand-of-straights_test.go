package main

import "testing"

func Test_isNStraightHand(t *testing.T) {
	tests := []struct {
		name string
		hand []int
		k    int
		want bool
	}{
		{
			hand: []int{1, 2, 2, 3, 3, 3, 4, 4, 5},
			k:    4,
			want: false,
		},

		{
			hand: []int{1, 2, 3, 4, 5},
			k:    4,
			want: false,
		},

		{
			hand: []int{1, 2, 3, 3, 4, 4, 5, 6},
			k:    4,
			want: true,
		},

		{
			hand: []int{3, 3, 2, 2, 1, 1},
			k:    3,
			want: true,
		},

		{
			hand: []int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11},
			k:    3,
			want: true,
		},

		{
			hand: []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
			k:    3,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNStraightHand(tt.hand, tt.k); got != tt.want {
				t.Errorf("isNStraightHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
