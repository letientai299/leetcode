package main

import "testing"

func Test_numPairsDivisibleBy60(t *testing.T) {
	tests := []struct {
		time []int
		want int
	}{
		{
			time: []int{30, 90, 150, 20, 80, 40},
			want: 5,
		},

		{
			time: []int{60, 60, 60},
			want: 3,
		},

		{
			time: []int{30, 20, 150, 100, 40},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := numPairsDivisibleBy60(tt.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}
