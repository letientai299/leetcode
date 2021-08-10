package main

import "testing"

func Test_makeFancyString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "aaabaaaa", want: "aabaa"},
		{s: "leeetcode", want: "leetcode"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := makeFancyString(tt.s); got != tt.want {
				t.Errorf("makeFancyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
