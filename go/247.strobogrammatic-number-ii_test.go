package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findStrobogrammatic(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{n: 3, want: []string{"111", "101", "181", "808", "818", "888", "609", "906", "619", "916", "689", "986"}},
		{n: 2, want: []string{"11", "88", "69", "96"}},
		{n: 0},
		{n: 1, want: []string{"1", "0", "8"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findStrobogrammatic(tt.n); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("findStrobogrammatic() = %v, want %v", got, tt.want)
			}
		})
	}
}
