package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generatePossibleNextMoves(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{s: "+-+++", want: []string{"+---+", "+-+--"}},
		{s: "+++++", want: []string{"--+++", "+--++", "++--+", "+++--"}},
		{s: "++", want: []string{"--"}},
		{s: "--", want: []string{}},
		{s: "", want: []string{}},
		{s: "+", want: []string{}},
		{s: "-", want: []string{}},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := generatePossibleNextMoves(tt.s)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func Benchmark_generatePossibleNextMoves(b *testing.B) {
	s := "+++++"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		generatePossibleNextMoves(s)
	}
}
