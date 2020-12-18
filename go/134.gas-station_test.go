package main

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	tests := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		{
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			want: -1,
		},

		{
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.gas, tt.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
