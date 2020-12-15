package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_simplifiedFractions(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			n: 4, want: []string{"1/2", "1/3", "1/4", "2/3", "3/4"},
		}, {
			n: 5, want: []string{"1/2", "1/3", "1/4", "1/5", "2/3", "2/5", "3/4", "3/5", "4/5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifiedFractions(tt.n); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("simplifiedFractions() = %v, want %v", got, tt.want)
			}
		})
	}
}
