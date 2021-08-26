package main

import "testing"

func Test_numberOfSubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "ababbbc", want: 3},
		// 1-2: 7 + 6 = 13
		// bbbb c -> 3
		// aba bbbb: 4+3+2+1 = 10
		{s: "abcabc", want: 10},
		{s: "aabaacabb", want: 18},
		{s: "abc", want: 1},
		{s: "aaaaa", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := numberOfSubstrings(tt.s); got != tt.want {
				t.Errorf("numberOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
