package main

import (
	"fmt"
	"testing"
)

func Test_addDigits(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := addDigits(tt.num)
			if got != tt.want {
				t.Errorf("addDigits(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
