package main

import "testing"

func Test_hIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{citations: []int{0, 1, 1, 1, 1}, want: 1},
		{citations: []int{0, 1, 2, 5, 6}, want: 2},
		{citations: []int{3, 0, 6, 1, 5}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
