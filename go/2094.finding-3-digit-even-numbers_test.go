package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findEvenNumbers(t *testing.T) {
	tests := []struct {
		digits []int
		want   []int
	}{
		{
			digits: []int{2, 2, 8, 8, 2},
			want:   []int{222, 228, 282, 288, 822, 828, 882},
		},
		{
			digits: []int{2, 1, 3, 0},
			want:   []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.digits), func(t *testing.T) {
			assert.Equalf(t, tt.want, findEvenNumbers(tt.digits), "findEvenNumbers(%v)", tt.digits)
		})
	}
}
