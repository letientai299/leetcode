package main

import (
	"fmt"
	"testing"
)

func Test_isThree(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{n: 1, want: false},
		{n: 2, want: false},
		{n: 3, want: false},
		{n: 5, want: false},
		{n: 6, want: false},
		{n: 12, want: false},
		{n: 16, want: false},
		{n: 13, want: false},
		{n: 17, want: false},

		{n: 4, want: true},
		{n: 9, want: true},
		{n: 49, want: true},
		{n: 121, want: true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			if got := isThree(tt.n); got != tt.want {
				t.Errorf("isThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
