package main

import "testing"

func Test_minStoneSum(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		k     int
		want  int
	}{
		{piles: []int{4, 3, 6, 7}, k: 3, want: 12},
		{piles: []int{5, 4, 9}, k: 2, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStoneSum(tt.piles, tt.k); got != tt.want {
				t.Errorf("minStoneSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
