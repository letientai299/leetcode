package main

import (
	"fmt"
	"testing"
)

func Test_getSum(t *testing.T) {
	tests := []struct {
		a int
		b int
	}{
		{1, 2},
		{-1, 2},
		{1000000, 20000},
		{-1000000, 20000},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.a, tt.b,
		)
		t.Run(testName, func(t *testing.T) {
			got := getSum(tt.a, tt.b)
			want := tt.a + tt.b
			if got != want {
				t.Errorf("getSum(%v, %v) = %v, want %v", tt.a, tt.b, got, want)
			}
		})
	}
}
