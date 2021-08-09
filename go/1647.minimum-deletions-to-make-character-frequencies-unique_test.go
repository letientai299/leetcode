package main

import "testing"

func Test_minDeletions(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "abcd", want: 3},
		{s: "aab", want: 0},
		{s: "aaabbbcc", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minDeletions(tt.s); got != tt.want {
				t.Errorf("minDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
