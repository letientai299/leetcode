package main

import "testing"

func Test_maxUniqueSplit(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "wwwzfvedwfvhsww", want: 11},
		// wwwzfvedwfvhsww
		// w ww z f v e d fv

		{s: "ababccc", want: 5},
		{s: "aa", want: 1},
		{s: "aba", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := maxUniqueSplit(tt.s); got != tt.want {
				t.Errorf("maxUniqueSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
