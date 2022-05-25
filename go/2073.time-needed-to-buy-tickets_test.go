package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_timeRequiredToBuy(t *testing.T) {
	tests := []struct {
		name    string
		tickets []int
		k       int
		want    int
	}{
		{
			tickets: []int{
				15, 66, 3, 47, 71, 27, 54, 43, 97, 34, 94, 33, 54, 26, 15, 52, 20, 71, 88, 42, 50, 6, 66, 88, 36, 99, 27, 82, 7, 72,
			},
			k:    18,
			want: 1457,
		},
		{
			tickets: []int{1, 1, 5, 2},
			k:       2,
			want:    9,
		},

		{
			tickets: []int{5, 1, 1, 1},
			k:       0,
			want:    8,
		},
		{
			tickets: []int{2, 3, 2},
			k:       2,
			want:    6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, timeRequiredToBuy(tt.tickets, tt.k), "timeRequiredToBuy(%v, %v)", tt.tickets, tt.k)
		})
	}
}
