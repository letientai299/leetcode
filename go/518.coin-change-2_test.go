package main

import "testing"

func Test_change(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		coins  []int
		want   int
	}{
		{amount: 0, coins: []int{7}, want: 1},
		{amount: 5, coins: []int{1, 2, 5}, want: 4},
		{amount: 5, coins: []int{1, 3, 4}, want: 3},
		// 4 1
		// 3 1 1
		// 1 1 1 1 1
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.amount, tt.coins); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
		})
	}
}
