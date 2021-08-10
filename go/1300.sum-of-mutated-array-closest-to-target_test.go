package main

import "testing"

func Test_findBestValue(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{arr: []int{60864, 25176, 27249, 21296, 20204}, target: 56803, want: 11361},
		{arr: []int{2, 3, 5}, target: 10, want: 5},
		{arr: []int{4, 3, 9}, target: 10, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBestValue(tt.arr, tt.target); got != tt.want {
				t.Errorf("findBestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
