package main

import "testing"

func Test_maxIceCream(t *testing.T) {
	tests := []struct {
		name  string
		costs []int
		coins int
		want  int
	}{
		{costs: []int{1, 3, 2, 4, 1}, coins: 7, want: 4},
		{costs: []int{10, 6, 8, 7, 7, 8}, coins: 5, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIceCream(tt.costs, tt.coins); got != tt.want {
				t.Errorf("maxIceCream() = %v, want %v", got, tt.want)
			}
		})
	}
}
