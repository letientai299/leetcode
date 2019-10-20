package main

import (
	"fmt"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{t: "abnc", s: "abc", want: true},
		{t: "abnc", s: "abd", want: false},
		{t: "abnc", s: "bc", want: true},
		{t: "abnc", s: "ac", want: true},
		{t: "abnc", s: "ca", want: false},
		{t: "abnc", s: "ca", want: false},
		{t: "abnc", s: "ca", want: false},
		{t: "abnca", s: "ca", want: true},
		{t: "abnca", s: "aca", want: true},
		{t: "abnca", s: "caa", want: false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.s, tt.t,
		)
		t.Run(testName, func(t *testing.T) {
			got := isSubsequence(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isSubsequence(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
