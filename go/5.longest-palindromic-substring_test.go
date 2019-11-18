package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{s: "a", want: []string{"a"}},
		{s: "", want: []string{""}},
		{s: "12313312312312aaaa123", want: []string{"aaaa", "1331"}},
		{s: "aaaa123", want: []string{"aaaa"}},
		{s: "babad", want: []string{"bab", "aba"}},
		{s: "bbad", want: []string{"bb"}},
		{s: "bbaba", want: []string{"bab"}},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			assert.Contains(t, tt.want, got)
		})
	}
}
