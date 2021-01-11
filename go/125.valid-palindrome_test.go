package main

import "testing"

func Test_isPalindrome_1(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"1", true},
		{"11", true},
		{"121", true},
		{"    ", true},
		{"3232131", false},
		{"A man, a plan, a canal: Panama", true},
		{"!A man, a plan, a canal: Panama", true},
		{"123A man, a plan, a canal : Panama 321", true},
		{"race a car", false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := isPalindrome_1(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
