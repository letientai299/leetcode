package main

import (
	"strconv"
	"testing"
)

func Test_isPalindrome_i(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{
			in:   123,
			want: false,
		},
		{
			in:   -123,
			want: false,
		},
		{
			in:   12320,
			want: false,
		},
		{
			in:   0,
			want: true,
		},
		{
			in:   11,
			want: true,
		},
		{
			in:   1221,
			want: true,
		},
		{
			in:   1100000011,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			if got := isPalindrome_i(tt.in); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
