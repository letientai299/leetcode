package main

import "testing"

func Test_canFormArray(t *testing.T) {
	tests := []struct {
		arr    []int
		pieces [][]int
		want   bool
	}{
		{
			arr:    []int{1, 2, 3},
			pieces: [][]int{{3, 3}, {1}},
			want:   false,
		},

		{
			arr:    []int{1, 2, 3},
			pieces: [][]int{{2, 3}, {1}},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canFormArray(tt.arr, tt.pieces); got != tt.want {
				t.Errorf("canFormArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
