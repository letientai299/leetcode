package main

import "testing"

func Test_containsPattern(t *testing.T) {
	tests := []struct {
		arr  []int
		m    int
		k    int
		want bool
	}{
		{
			arr:  []int{1, 1, 2, 2, 2, 2},
			m:    1,
			k:    3,
			want: true,
		},

		{
			arr:  []int{2, 2},
			m:    1,
			k:    2,
			want: true,
		},

		{
			arr:
			[]int{1, 2, 1, 2, 1, 3},
			m:    2,
			k:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := containsPattern(tt.arr, tt.m, tt.k); got != tt.want {
				t.Errorf("containsPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
