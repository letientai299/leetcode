package main

import "testing"

func Test_closeStrings(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{a: "aaabbbbccddeeeeefffff", b: "aaaaabbcccdddeeeeffff", want: false},
		{a: "abcd", b: "bca", want: false},
		{a: "aab", b: "bca", want: false},
		{a: "abc", b: "dcs", want: false},

		{a: "abc", b: "bca", want: true},
		{a: "cabbba", b: "abbccc", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closeStrings(tt.a, tt.b); got != tt.want {
				t.Errorf("closeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
