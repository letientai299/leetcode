package main

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	tests := []struct {
		str  string
		sub  string
		want int
	}{
		{
			"hello", "llx", -1,
		},
		{
			"hello", "ll", 2,
		},
		{
			"aaabb", "bb", 3,
		},
		{
			"", "bb", -1,
		},
		{
			"", "", 0,
		},
		{
			"aaabb", "ab", 2,
		},
		{
			"aaabb", "aa", 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s | %s", tt.str, tt.sub), func(t *testing.T) {
			got := strStr(tt.str, tt.sub)
			if got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
