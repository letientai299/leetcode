package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minBitFlips(t *testing.T) {
	tests := []struct {
		name  string
		start int
		goal  int
		want  int
	}{
		{
			start: 1,
			goal:  2,
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minBitFlips(tt.start, tt.goal), "minBitFlips(%v, %v)", tt.start, tt.goal)
		})
	}
}
