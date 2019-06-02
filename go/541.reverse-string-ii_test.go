package main

import (
	"fmt"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{"abcdefg", 2, "bacdfeg"},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.s, tt.k,
		)
		t.Run(testName, func(t *testing.T) {
			got := reverseStr(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("reverseStr(%v, %v) = %v, want %v", tt.s, tt.k, got, tt.want)
			}
		})
	}
}
