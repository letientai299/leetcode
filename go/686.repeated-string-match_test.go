package main

import (
	"fmt"
	"testing"
)

func Test_repeatedStringMatch(t *testing.T) {
	tests := []struct {
		A    string
		B    string
		want int
	}{
		{
			A:    "abcd",
			B:    "bcdab",
			want: 2,
		},

		{
			A:    "aaaaaaaaaaaaaaaaaaaaaab",
			B:    "ba",
			want: 2,
		},

		{
			A:    "ab",
			B:    "ababab",
			want: 3,
		},

		{
			A:    "abcd",
			B:    "xabcd",
			want: -1,
		},

		{
			A:    "ab",
			B:    "bab",
			want: 2,
		},

		{
			A:    "ab",
			B:    "aab",
			want: -1,
		},

		{
			A:    "ab",
			B:    "abb",
			want: -1,
		},

		{
			A:    "ab",
			B:    "abab",
			want: 2,
		},

		{
			A:    "ab",
			B:    "abaabab",
			want: -1,
		},

		{
			A:    "abcd",
			B:    "abcd",
			want: 1,
		},

		{
			A:    "abcd",
			B:    "cdabcdabcdab",
			want: 4,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.A, tt.B,
		)
		t.Run(testName, func(t *testing.T) {
			got := repeatedStringMatch(tt.A, tt.B)
			if got != tt.want {
				t.Errorf("repeatedStringMatch(%v, %v) = %v, want %v", tt.A, tt.B, got, tt.want)
			}
		})
	}
}
