package main

import (
	"fmt"
	"math"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		str  string
		want int
	}{
		{str: "-6147483648", want: math.MinInt32},
		{str: "-91283472332", want: math.MinInt32},
		{str: "      -9000000000000000000000", want: math.MinInt32},
		{str: "      9000000000000000000000", want: math.MaxInt32},
		{str: "      +9000000000000000000000", want: math.MaxInt32},
		{str: "      2147483647", want: 2147483647},
		{str: "      0000123", want: 123},
		{str: "      -0000123", want: -123},
		{str: "      +123", want: 123},
		{str: "      -123aaaa", want: -123},
		{str: "      -123     ", want: -123},
		{str: "123", want: 123},
		{str: "-123", want: -123},
		{str: "", want: 0},
		{str: "      ", want: 0},
		{str: "-+123", want: 0},
		{str: "a123", want: 0},
		{str: "a", want: 0},
		{str: "    a", want: 0},
		{str: " \n   a", want: 0},
		{str: "-+123", want: 0},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.str,
		)
		t.Run(testName, func(t *testing.T) {
			got := myAtoi(tt.str)
			if got != tt.want {
				t.Errorf("myAtoi(%v) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}
