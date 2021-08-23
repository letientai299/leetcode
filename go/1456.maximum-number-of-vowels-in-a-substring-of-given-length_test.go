package main

import "testing"

func Test_maxVowels(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{s: "leetcode", k: 3, want: 2},
		{s: "rhythms", k: 4, want: 0},
		{s: "abciiidef", k: 3, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := maxVowels(tt.s, tt.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
