package main

import "testing"

func Test_countTriplets(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			arr:  []int{2, 3, 1, 6, 7},
			want: 4,
		},

		{
			arr:  []int{1, 1, 1, 1, 1},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriplets(tt.arr); got != tt.want {
				t.Errorf("countTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
