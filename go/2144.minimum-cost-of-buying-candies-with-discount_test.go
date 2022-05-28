package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumCost(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			cost: []int{1, 2, 3},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumCost(tt.cost), "minimumCost(%v)", tt.cost)
		})
	}
}
