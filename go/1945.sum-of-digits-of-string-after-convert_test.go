package main

import "testing"

func Test_getLucky(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{s: "fleyctuuajsr", k: 5, want: 8},
		{s: "leetcode", k: 2, want: 6},
		{s: "iiii", k: 1, want: 36},
		{s: "zbax", k: 2, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLucky(tt.s, tt.k); got != tt.want {
				t.Errorf("getLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
