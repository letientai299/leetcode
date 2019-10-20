package main

import (
	"fmt"
	"testing"
)

func Test_isToeplitzMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   bool
	}{
		{
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			want: true,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.matrix,
		)
		t.Run(testName, func(t *testing.T) {
			got := isToeplitzMatrix(tt.matrix)
			if got != tt.want {
				t.Errorf("isToeplitzMatrix(%v) = %v, want %v", tt.matrix, got, tt.want)
			}
		})
	}
}
