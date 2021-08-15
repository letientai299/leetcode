package main

import "testing"

func Test_mctFromLeafValues(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			arr:  []int{8, 10, 6, 5, 3, 2, 4, 5, 13, 15, 13, 15, 4, 3, 11, 1, 14, 9, 9, 4},
			want: 1652,
		},
		{arr: []int{4, 11}, want: 44},
		{arr: []int{6, 2, 4}, want: 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mctFromLeafValues(tt.arr); got != tt.want {
				t.Errorf("mctFromLeafValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
