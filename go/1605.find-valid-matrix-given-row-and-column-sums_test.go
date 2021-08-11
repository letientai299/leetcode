package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_restoreMatrix(t *testing.T) {
	tests := []struct {
		name   string
		rowSum []int
		colSum []int
	}{
		{
			rowSum: []int{14, 9},
			colSum: []int{6, 9, 8},
		},

		{
			rowSum: []int{5, 7, 10},
			colSum: []int{8, 6, 8},
		},

		{
			rowSum: []int{3, 8},
			colSum: []int{4, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := len(tt.rowSum)
			rows := make([]int, m)
			copy(rows, tt.rowSum)

			n := len(tt.colSum)
			cols := make([]int, n)
			copy(cols, tt.colSum)

			got := restoreMatrix(rows, cols)
			for y := 0; y < m; y++ {
				for x := 0; x < n; x++ {
					tt.colSum[x] -= got[y][x]
					tt.rowSum[y] -= got[y][x]
				}
			}

			for _, v := range tt.colSum {
				assert.Equal(t, 0, v, got)
			}

			for _, v := range tt.rowSum {
				assert.Equal(t, 0, v, got)
			}
		})
	}
}
