package main

import (
	"fmt"
	"testing"
)

func Test_findLUSlength521(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{
			a:    "aba",
			b:    "cdc",
			want: 3,
		},

		{
			a:    "aba",
			b:    "adc",
			want: 3,
		},

		{
			a:    "aba",
			b:    "abc",
			want: 3,
		},

		{
			a:    "aba",
			b:    "cdc",
			want: 3,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.a, tt.b,
		)
		t.Run(testName, func(t *testing.T) {
			got := findLUSlength521(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("findLUSlength(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
