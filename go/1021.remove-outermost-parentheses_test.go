package main

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "(()())", want: "()()"},
		{s: "()()", want: ""},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := removeOuterParentheses(tt.s); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
