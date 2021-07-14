package main

import "testing"

func Test_totalMoney(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 1, want: 1},
		{n: 4, want: 10},
		{n: 8, want: 30},
		{n: 10, want: 37},
		{n: 20, want: 96},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalMoney(tt.n); got != tt.want {
				t.Errorf("totalMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
