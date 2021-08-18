package main

import "testing"

func Test_scoreOfParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "(()(())())()", want: 9},
		{s: "(()(()))", want: 6},
		{s: "(()())", want: 4},
		{s: "(())", want: 2},
		{s: "()", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := scoreOfParentheses(tt.s); got != tt.want {
				t.Errorf("scoreOfParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
