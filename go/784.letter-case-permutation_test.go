package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCasePermutation(t *testing.T) {
	tests := []struct {
		S    string
		want []string
	}{
		{S: "ab", want: []string{"ab", "Ab", "aB", "AB"}},
		{S: "a1b", want: []string{"a1b", "A1b", "a1B", "A1B"}},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.S,
		)
		t.Run(testName, func(t *testing.T) {
			got := letterCasePermutation(tt.S)
			assert.ElementsMatch(t, got, tt.want, "letterCasePermutation(%v) = %v, want %v", tt.S, got, tt.want)
		})
	}
}

func Benchmark_letterCasePermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = letterCasePermutation("a1b2c")
	}
}
