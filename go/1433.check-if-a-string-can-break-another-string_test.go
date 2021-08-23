package main

import "testing"

func Test_checkIfCanBreak(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{s1: "abc", s2: "xya", want: true},
		{s1: "eba", s2: "cda", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfCanBreak(tt.s1, tt.s2); got != tt.want {
				t.Errorf("checkIfCanBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
