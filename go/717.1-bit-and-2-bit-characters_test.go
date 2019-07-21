package main

import (
	"fmt"
	"testing"
)

func Test_isOneBitCharacter(t *testing.T) {
	tests := []struct {
		bits []int
		want bool
	}{
		{[]int{1, 0, 0}, true},
		{[]int{1, 1, 0, 0}, true},
		{[]int{1, 1, 0}, true},
		{[]int{1, 1, 1, 0}, false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.bits,
		)
		t.Run(testName, func(t *testing.T) {
			got := isOneBitCharacter(tt.bits)
			if got != tt.want {
				t.Errorf("isOneBitCharacter(%v) = %v, want %v", tt.bits, got, tt.want)
			}
		})
	}
}
