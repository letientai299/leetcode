package main

import "testing"

func Test_numOfSubarrays(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		k         int
		threshold int
		want      int
	}{
		{
			arr:       []int{2, 2, 2, 2, 5, 5, 5, 8},
			k:         3,
			threshold: 4,
			want:      3,
		},

		{
			arr:       []int{1, 1, 1, 1, 1},
			k:         1,
			threshold: 0,
			want:      5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfSubarrays(tt.arr, tt.k, tt.threshold); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
