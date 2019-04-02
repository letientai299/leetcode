package main

import (
	"fmt"
	"testing"
)

func Test_countSegments(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"Hello, world!", 2},
		{"   Hello, world!", 2},
		{"   Hello, world!   ", 2},
		{"   Hello,    world!   ", 2},
		{"   Hello    ,    world!   ", 3},
		{"Hong biết thgian đi đâu mất", 6},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := countSegments(tt.s)
			if got != tt.want {
				t.Errorf("countSegments(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
