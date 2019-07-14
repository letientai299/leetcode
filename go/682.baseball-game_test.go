package main

import (
	"fmt"
	"testing"
)

func Test_calPoints(t *testing.T) {
	tests := []struct {
		ops  []string
		want int
	}{

		{
			ops:  []string{"5", "2", "C", "D", "+"},
			want: 30,
		},

		{
			ops:  []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			want: 27,
		},
	}

	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.ops,
		)
		t.Run(testName, func(t *testing.T) {
			got := calPoints(tt.ops)
			if got != tt.want {
				t.Errorf("calPoints(%v) = %v, want %v", tt.ops, got, tt.want)
			}
		})
	}
}
