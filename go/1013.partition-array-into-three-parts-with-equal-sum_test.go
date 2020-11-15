package main

import "testing"

func Test_canThreePartsEqualSum(t *testing.T) {
	tests := []struct {
		a    []int
		want bool
	}{
		{
			a:    []int{1, -1, 1, -1},
			want: false,
		},

		{
			a:    []int{10, -10, 10, -10, 10, -10, 10, -10},
			want: true,
		},

		{
			a:    []int{6, 1, 1, 13, -1, 0, -10, 20},
			want: false,
		},

		{
			a:    []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canThreePartsEqualSum(tt.a); got != tt.want {
				t.Errorf("canThreePartsEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
