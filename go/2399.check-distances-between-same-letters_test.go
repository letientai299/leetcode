package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkDistances(t *testing.T) {
	tests := []struct {
		s        string
		distance []int
		want     bool
	}{
		{
			s: "abaccb",
			distance: []int{
				1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			assert.Equalf(t, tt.want, checkDistances(tt.s, tt.distance), "checkDistances(%v, %v)", tt.s, tt.distance)
		})
	}
}
