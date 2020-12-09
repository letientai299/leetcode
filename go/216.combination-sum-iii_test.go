package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum3(t *testing.T) {
	tests := []struct {
		name string
		k    int
		n    int
		want [][]int
	}{
		{n: 1, k: 4, want: [][]int{}},
		{n: 9, k: 3, want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{n: 7, k: 3, want: [][]int{{1, 2, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.k, tt.n); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
