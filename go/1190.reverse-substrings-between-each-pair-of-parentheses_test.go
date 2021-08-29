package main

import "testing"

func Test_reverseParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "(abcd)", want: "dcba"},
		{s: "abcd", want: "abcd"},
		{s: "a(bc)d", want: "acbd"},
		{s: "a(b)cd", want: "abcd"},
		{s: "a((bc)d)", want: "adbc"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := reverseParentheses(tt.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
