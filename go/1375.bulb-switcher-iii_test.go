package main

import "testing"

func Test_numTimesAllBlue(t *testing.T) {
	tests := []struct {
		name  string
		light []int
		want  int
	}{
		// 1 1 1 1 0
		{light: []int{3, 2, 4, 1, 5}, want: 2},
		{light: []int{4, 1, 2, 3}, want: 1},
		{light: []int{2, 1, 3, 5, 4}, want: 3},
		{light: []int{1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTimesAllBlue(tt.light); got != tt.want {
				t.Errorf("numTimesAllBlue() = %v, want %v", got, tt.want)
			}
		})
	}
}
