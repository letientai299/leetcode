package main

import "testing"

func Test_countAndSay(t *testing.T) {
	tests := []struct {
		n int
		s string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
		{7, "13112221"},
		{8, "1113213211"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := countAndSay(tt.n); got != tt.s {
				t.Errorf("countAndSay() = %v, want %v", got, tt.s)
			}
		})
	}
}
