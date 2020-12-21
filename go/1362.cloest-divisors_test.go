package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_closestDivisors(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want []int
	}{
		{num: 10, want: []int{3, 4}},
		{num: 9, want: []int{2, 5}},
		{num: 4, want: []int{2, 3}},
		{num: 123, want: []int{5, 25}},
		{num: 722900699, want: []int{25140, 28755}},
		{num: 999, want: []int{25, 40}},
		{num: 8, want: []int{3, 3}},
		{num: 688427155, want: []int{2409, 285773}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestDivisors(tt.num); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("closestDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
