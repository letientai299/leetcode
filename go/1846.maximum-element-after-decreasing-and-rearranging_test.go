package main

import "testing"

func Test_maximumElementAfterDecrementingAndRearranging(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			arr:  []int{2, 2, 1, 2, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumElementAfterDecrementingAndRearranging(tt.arr); got != tt.want {
				t.Errorf("maximumElementAfterDecrementingAndRearranging() = %v, want %v", got, tt.want)
			}
		})
	}
}
