package main

import (
	"reflect"
	"testing"
)

func Test_numOfBurgers(t *testing.T) {
	tests := []struct {
		name         string
		tomatoSlices int
		cheeseSlices int
		want         []int
	}{
		{
			tomatoSlices: 2385088,
			cheeseSlices: 164763,
			want:         nil,
		},

		{
			tomatoSlices: 0,
			cheeseSlices: 0,
			want:         []int{0, 0},
		},

		{
			tomatoSlices: 16,
			cheeseSlices: 7,
			want:         []int{1, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfBurgers(tt.tomatoSlices, tt.cheeseSlices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numOfBurgers() = %v, want %v", got, tt.want)
			}
		})
	}
}
