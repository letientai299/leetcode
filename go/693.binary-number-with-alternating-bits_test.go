package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_hasAlternatingBits(t *testing.T) {
	tests := []struct {
		bits string
		want bool
	}{
		{"0", true},
		{"1", true},
		{"10", true},
		{"11", false},
		{"1001", false},
		{"10101", true},
		{"1011", false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.bits,
		)
		t.Run(testName, func(t *testing.T) {
			n, _ := strconv.ParseInt(tt.bits, 2, 32)
			got := hasAlternatingBits(int(n))
			if got != tt.want {
				t.Errorf("hasAlternatingBits(%v) = %v, want %v", n, got, tt.want)
			}
		})
	}
}
