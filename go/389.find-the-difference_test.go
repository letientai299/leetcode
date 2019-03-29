package main

import (
	"fmt"
	"testing"
)

func Test_findTheDifference(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want byte
	}{
		{s: "abc", t: "abcd", want: 'd'},
		{s: "abc", t: "abca", want: 'a'},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.s, tt.t,
		)
		t.Run(testName, func(t *testing.T) {
			got := findTheDifference(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("findTheDifference(%v, %v) = %v, want %s", tt.s, tt.t, got, string(tt.want))
			}
		})
	}
}
