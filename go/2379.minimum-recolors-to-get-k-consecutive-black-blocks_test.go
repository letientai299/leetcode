package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumRecolors(t *testing.T) {
	tests := []struct {
		blocks string
		k      int
		want   int
	}{
		{blocks: "WBBWWWWBBWWBBBBWWBBWWBBBWWBBBWWWBWBWW", k: 15, want: 6},
		{blocks: "WBBWWBBWBW", k: 7, want: 3},
		{blocks: "WBWBBBW", k: 2, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.blocks, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumRecolors(tt.blocks, tt.k), "minimumRecolors(%v, %v)", tt.blocks, tt.k)
		})
	}
}
