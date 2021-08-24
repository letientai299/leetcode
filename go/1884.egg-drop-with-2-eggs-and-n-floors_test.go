package main

import "testing"

func Test_twoEggDrop(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 11, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoEggDrop(tt.n); got != tt.want {
				t.Errorf("twoEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}
