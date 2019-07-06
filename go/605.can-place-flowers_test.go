package main

import (
	"fmt"
	"testing"
)

func Test_canPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{
			[]int{1, 0, 0, 0, 1},
			1,
			true,
		},

		{
			[]int{1, 0, 0, 0, 1},
			2,
			false,
		},

		{
			[]int{0},
			1,
			true,
		},

		{
			[]int{0},
			0,
			true,
		},

		{
			[]int{1, 0},
			0,
			true,
		},

		{
			[]int{1, 0},
			1,
			false,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.flowerbed, tt.n,
		)
		t.Run(testName, func(t *testing.T) {
			got := canPlaceFlowers(tt.flowerbed, tt.n)
			if got != tt.want {
				t.Errorf("canPlaceFlowers(%v, %v) = %v, want %v", tt.flowerbed, tt.n, got, tt.want)
			}
		})
	}
}
