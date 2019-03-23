package main

import (
	"fmt"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	tests := []struct {
		want int
		s    string
	}{
		{
			1, "A",
		},
		{
			2, "B",
		},
		{
			701, "ZY",
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := titleToNumber(tt.s)
			if got != tt.want {
				t.Errorf("titleToNumber(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
