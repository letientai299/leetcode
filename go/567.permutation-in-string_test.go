package main

import "testing"

func Test_checkInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{s1: "ab", s2: "ccab", want: true},
		{s1: "abb", s2: "bcababba", want: true},
		{s1: "ab", s2: "bacc", want: true},
		{s1: "ab", s2: "ccadb", want: false},
		{s1: "ab", s2: "ccba", want: true},
		{s1: "ab", s2: "cba", want: true},
		{s1: "ab", s2: "eidboaoo", want: false},
		{s1: "ab", s2: "eidbaooo", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.s1, tt.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
