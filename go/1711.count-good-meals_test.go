package main

import "testing"

func Test_countPairs(t *testing.T) {
	tests := []struct {
		name          string
		deliciousness []int
		want          int
	}{
		{deliciousness: []int{1048576, 1048576}, want: 1},
		{deliciousness: []int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234}, want: 12},
		{deliciousness: []int{1, 1, 1, 3, 3, 3, 7}, want: 15},
		{deliciousness: []int{1, 3, 5, 7, 9}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.deliciousness); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
