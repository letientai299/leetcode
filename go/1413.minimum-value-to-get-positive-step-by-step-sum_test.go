package main

import (
	"testing"
)

func Test_minStartValue(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{
			a:    []int{-3, 2, -3, 4, 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minStartValue(tt.a); got != tt.want {
				t.Errorf("minStartValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
