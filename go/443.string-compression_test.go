package main

import (
	"fmt"
	"testing"
)

func Test_compress(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"a", 1},
		{"aa", 2},
		{"aaa", 2},
		{"aaaaaa", 2},
		{"aabbccc", 6},
		{"aabbccca", 7},
		{"aabbccca", 7},
		{"baaaaaaaaaaaabbccca", 9},
		{"aaaaaaaaaabb", 5},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			in := []byte(tt.s)
			got := compress(in)
			fmt.Println(string(in[0:got]))
			if got != tt.want {
				t.Errorf("compress(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
