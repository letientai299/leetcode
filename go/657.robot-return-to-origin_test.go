package main

import (
	"fmt"
	"testing"
)

func Test_judgeCircle(t *testing.T) {
	tests := []struct {
		moves string
		want  bool
	}{
		{"UD", true},
		{"UDLR", true},
		{"UDRL", true},
		{"UUDRL", false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.moves,
		)
		t.Run(testName, func(t *testing.T) {
			got := judgeCircle(tt.moves)
			if got != tt.want {
				t.Errorf("judgeCircle(%v) = %v, want %v", tt.moves, got, tt.want)
			}
		})
	}
}
