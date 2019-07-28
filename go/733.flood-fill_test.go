package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	tests := []struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
		want     [][]int
	}{
		{
			image: [][]int{

				{0, 0, 0},
				{0, 1, 1},
			},

			sr:       1,
			sc:       1,
			newColor: 1,
			want: [][]int{
				{0, 0, 0},
				{0, 1, 1},
			},
		},

		{
			image: [][]int{

				{3, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},

			sr:       0,
			sc:       0,
			newColor: 2,
			want: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
		},

		{
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},

			sr:       1,
			sc:       1,
			newColor: 2,
			want: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},

		{
			image: [][]int{

				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},

			sr:       0,
			sc:       0,
			newColor: 2,
			want: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v, %v, %v",
			tt.image, tt.sr, tt.sc, tt.newColor,
		)
		t.Run(testName, func(t *testing.T) {
			if got := floodFill(tt.image, tt.sr, tt.sc, tt.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill(%v, %v, %v, %v) = %v, want %v", tt.image, tt.sr, tt.sc, tt.newColor, got, tt.want)
			}
		})
	}
}
