package main

import "testing"

func Test_isValid20(t *testing.T) {
	tests := []struct {
		args string
		want bool
	}{
		{"{}", true},
		{"", true},
		{"{()}", true},
		{"{}", true},
		{"}", false},
		{"(]", false},
		{"(", false},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := isValid20(tt.args); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
