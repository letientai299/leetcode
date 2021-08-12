package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_constructArray(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{n: 4},
		{n: 10},
		{n: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k := 1; k < tt.n; k++ {
				got := constructArray(tt.n, k)
				m := make(map[int]int)
				for i := 0; i < tt.n-1; i++ {
					v := got[i] - got[i+1]
					if v < 0 {
						v = -v
					}
					m[v]++
				}
				assert.Lenf(t, m, k, "got=%v, m=%v", got, m)
			}
		})
	}
}
