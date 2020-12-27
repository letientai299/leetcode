package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_diffWaysToCompute(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{
			input: "2*3-4*5",
			want:  []int{-34, -14, -10, -10, 10},
		},

		{
			input: "2-1-1",
			want:  []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := diffWaysToCompute(tt.input); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
