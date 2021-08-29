package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{s: "deeedbbcccbdaa", k: 3, want: "aa"},
		{s: "", want: ""},
		{s: "abcd", k: 2, want: "abcd"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := removeDuplicates(tt.s, tt.k); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
