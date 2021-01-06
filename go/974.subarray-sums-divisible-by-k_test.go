package main

import "testing"

func Test_subarraysDivByK(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		k    int
		want int
	}{
		{a: []int{4, 5, 0, -2, -3, 1}, k: 5, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(tt.a, tt.k); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
