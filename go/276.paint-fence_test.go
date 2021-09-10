package main

import (
	"fmt"
	"testing"
)

func Test_numWays276(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		// aab, aba, abb, bab, bba, baa
		{n: 3, k: 2, want: 6},

		{n: 2, k: 1, want: 1},

		{n: 1, k: 2, want: 2},

		// __aa: _(b|c)aa
		// __bb: _(
		// __cc:
		{n: 4, k: 3, want: 66},
		// aab, aac, bba, bbc, cca, ccb
		// abb, aba, abc
		// aca, acb, acc
		// baa, bab, bac
		// bca, bcb, bcc
		// caa, cab, cac
		// cba, cbb, cbc
		{n: 3, k: 3, want: 24},

		// aa, ab, ac,
		// bb, ba, bc
		// cc, ca, cb
		{n: 2, k: 3, want: 9},
		{n: 1, k: 3, want: 3},

		// aaba, abaa, abba, baba, bbaa, baab
		// aabb, abab, babb, bbab,
		{n: 4, k: 2, want: 10},
		// aa, ab, bb, ba
		{n: 2, k: 2, want: 4},
		{n: 3, k: 1, want: 0},
		{n: 1, k: 1, want: 1},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.n, tt.k,
		)
		t.Run(testName, func(t *testing.T) {
			got := numWays276(tt.n, tt.k)
			if got != tt.want {
				t.Errorf("numWays276(%v, %v) = %v, want %v", tt.n, tt.k, got, tt.want)
			}
		})
	}
}
