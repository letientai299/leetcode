package main

import (
	"fmt"
	"testing"
)

func Test_checkRecord(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"PPALLP", true},
		{"PPALLL", false},
		{"LALL", true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := checkRecord(tt.s)
			if got != tt.want {
				t.Errorf("checkRecord(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
