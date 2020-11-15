package main

import "testing"

func Test_restoreString(t *testing.T) {
	tests := []struct {
		s       string
		indices []int
		want    string
	}{
		{
			s:       "abcd",
			indices: []int{1, 3, 0, 2},
			want:    "cadb",
		},

		{
			s:       "abc",
			indices: []int{0, 2, 1},
			want:    "acb",
		},

		{
			s:       "abc",
			indices: []int{1, 2, 0},
			want:    "cab",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := restoreString(tt.s, tt.indices); got != tt.want {
				t.Errorf("restoreString() = %v, want %v", got, tt.want)
			}
		})
	}
}
