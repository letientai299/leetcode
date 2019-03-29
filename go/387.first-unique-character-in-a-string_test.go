package main

import (
	"fmt"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"", -1},
		{"aabb", -1},
		{"aab", 2},
		{"leetcode", 0},
		{"loveleetcode", 2},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := firstUniqChar(tt.s)
			if got != tt.want {
				t.Errorf("firstUniqChar(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
