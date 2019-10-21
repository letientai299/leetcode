package main

import (
	"fmt"
	"testing"
)

func Test_rotateString(t *testing.T) {
	tests := []struct {
		A    string
		B    string
		want bool
	}{
		{},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.A, tt.B,
		)
		t.Run(testName, func(t *testing.T) {
			got := rotateString(tt.A, tt.B)
			if got != tt.want {
				t.Errorf("rotateString(%v, %v) = %v, want %v", tt.A, tt.B, got, tt.want)
			}
		})
	}
}
