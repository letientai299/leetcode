package main

import "testing"

func Test_stringShift(t *testing.T) {
	tests := []struct {
		s     string
		shift [][]int
		want  string
	}{

		{

			s: "wpdhhcj",
			shift: [][]int{
				{0, 7}, {1, 7}, {1, 0}, {1, 3}, {0, 3}, {0, 6}, {1, 2},
			},
			want: "hcjwpdh",
		},

		{
			s:     "abc",
			shift: [][]int{{0, 1}, {1, 8}},
			want:  "cab",
		},

		{
			s:     "abc",
			shift: [][]int{},
			want:  "abc",
		},

		{
			s:     "abc",
			shift: [][]int{{0, 2}, {1, 1}},
			want:  "bca",
		},

		{
			s:     "abc",
			shift: [][]int{{0, 1}, {1, 1}},
			want:  "abc",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := stringShift(tt.s, tt.shift); got != tt.want {
				t.Errorf("stringShift() = %v, want %v", got, tt.want)
			}
		})
	}
}
