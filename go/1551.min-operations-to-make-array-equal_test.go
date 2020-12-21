package main

import "testing"

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 3, want: 2}, // 1 3 5
		{n: 4, want: 4}, // 1 3 5 7
		{n: 5, want: 6}, // 1 3 5 7 9
		{n: 6, want: 9}, // 1 3 5 7 9 11
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.n); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
