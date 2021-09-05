package main

import "testing"

func Test_getWinner(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want int
	}{
		{
			arr:  []int{1, 25, 35, 42, 68, 70},
			k:    1,
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWinner(tt.arr, tt.k); got != tt.want {
				t.Errorf("getWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
