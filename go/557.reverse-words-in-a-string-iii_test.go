package main

import (
	"fmt"
	"testing"
)

func Test_reverseWords557(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "",
			want: "",
		},

		{
			s:    "Let's take LeetCode contest",
			want: "s'teL ekat edoCteeL tsetnoc",
		},

		{
			s:    "hello world",
			want: "olleh dlrow",
		},
		{
			s:    "a b c",
			want: "a b c",
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := reverseWords557(tt.s)
			if got != tt.want {
				t.Errorf("reverseWords(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
