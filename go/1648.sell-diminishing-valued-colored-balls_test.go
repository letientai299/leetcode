package main

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name      string
		inventory []int
		orders    int
		want      int
	}{
		{inventory: []int{15, 13, 89, 1, 22, 11, 90, 13, 54, 100}, orders: 249, want: 13711},
		{inventory: []int{73, 98, 35, 100}, orders: 178, want: 11120},
		{inventory: []int{1000, 1000}, orders: 2, want: 2000},
		{inventory: []int{2, 8, 4, 10, 6}, orders: 20, want: 110},
		{inventory: []int{3, 5}, orders: 6, want: 19},
		{inventory: []int{1000000000}, orders: 1000000000, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.inventory, tt.orders); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
