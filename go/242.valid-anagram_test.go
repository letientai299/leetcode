package main

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{
			"anagram",
			"nagaram",
			true,
		},
		{
			"nagram",
			"nagaram",
			false,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.s, tt.t,
		)
		t.Run(testName, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isAnagram(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
