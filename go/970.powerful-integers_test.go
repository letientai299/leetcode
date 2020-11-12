package main

import (
	"reflect"
	"testing"
)

func Test_powerfulIntegers(t *testing.T) {
	tests := []struct {
		x     int
		y     int
		bound int
		want  []int
	}{
		{
			x: 2, y: 1, bound: 10, want: []int{2, 3, 5, 9},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := powerfulIntegers(tt.x, tt.y, tt.bound); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("powerfulIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
