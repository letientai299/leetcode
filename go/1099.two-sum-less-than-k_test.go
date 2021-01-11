package main

import "testing"

func Test_twoSumLessThanK(t *testing.T) {
	tests := []struct {
		a    []int
		k    int
		want int
	}{
		{
			a:    []int{34, 23, 1, 24, 75, 33, 54, 8},
			k:    60,
			want: 58,
		},

		{
			a:    []int{34, 23, 1, 24, 75, 33, 54, 8},
			k:    1,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := twoSumLessThanK(tt.a, tt.k); got != tt.want {
				t.Errorf("twoSumLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
