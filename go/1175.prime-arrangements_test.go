package main

import "testing"

func Test_numPrimeArrangements(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 100, want: 682289015},
		{n: 5, want: 12},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := numPrimeArrangements(tt.n); got != tt.want {
				t.Errorf("numPrimeArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
