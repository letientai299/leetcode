package main

import "testing"

func Test_countGoodNumbers(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int
	}{
		{n: 4, want: 400},
		{n: 15, want: 399999958},
		{n: 1, want: 5},
		{n: 50, want: 564908303},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodNumbers(tt.n); got != tt.want {
				t.Errorf("countGoodNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
