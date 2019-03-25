package main

import (
	"fmt"
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"egg", "tee", true},
		{"eggegg", "teetee", true},
		{"eggegg", "teetcc", false},
		{"", "teetcc", false},
		{"", "", true},
		{"x", "", false},
		{"egg", "gge", false},
		{"1", "2", true},
		{"x", "y", true},
		{"abc", "xyz", true},
		{"xx", "yy", true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.s, tt.t,
		)
		t.Run(testName, func(t *testing.T) {
			got := isIsomorphic(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isIsomorphic(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
