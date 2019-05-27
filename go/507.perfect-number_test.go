package main

import (
	"fmt"
	"testing"
)

func Test_checkPerfectNumber(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{36, false},
		{28, true},
		{6, true},
		{0, false},
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, false},
		{7, false},
		{8, false},
		{9, false},
		{10, false},
		{11, false},
		{12, false},
		{13, false},
		{14, false},
		{15, false},
		{16, false},
		{17, false},
		{18, false},
		{19, false},
		{20, false},
		{21, false},
		{22, false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := checkPerfectNumber(tt.num)
			if got != tt.want {
				t.Errorf("checkPerfectNumber(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
