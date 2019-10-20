package main

import (
	"fmt"
	"testing"
)

func Test_countPrimeSetBits(t *testing.T) {
	tests := []struct {
		L    int
		R    int
		want int
	}{
		{},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.L, tt.R,
		)
		t.Run(testName, func(t *testing.T) {
			got := countPrimeSetBits(tt.L, tt.R)
			if got != tt.want {
				t.Errorf("countPrimeSetBits(%v, %v) = %v, want %v", tt.L, tt.R, got, tt.want)
			}
		})
	}
}
