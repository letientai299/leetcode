package main

import (
	"fmt"
	"testing"
)

func Test_isStrobogrammatic(t *testing.T) {
	tests := []struct {
		num  string
		want bool
	}{
		{num: "8", want: true},
		{num: "0", want: true},
		{num: "88", want: true},
		{num: "808", want: true},
		{num: "808", want: true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := isStrobogrammatic(tt.num)
			if got != tt.want {
				t.Errorf("isStrobogrammatic(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
