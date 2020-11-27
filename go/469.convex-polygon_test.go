package main

import "testing"

func Test_isConvex(t *testing.T) {
	tests := []struct {
		ps   [][]int
		want bool
	}{
		{
			ps:   [][]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {-2, 1}, {-2, 0}, {-1, 0}},
			want: true,
		},

		{
			ps:   [][]int{{0, 0}, {1, 0}, {1, 1}, {-1, 1}, {-1, 0}},
			want: true,
		},

		{
			ps:   [][]int{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
			want: true,
		},

		{
			ps:   [][]int{{0, 0}, {0, 10}, {10, 10}, {10, 0}, {5, 5}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isConvex(tt.ps); got != tt.want {
				t.Errorf("isConvex() = %v, want %v", got, tt.want)
			}
		})
	}
}
