package main

import "testing"

func Test_divide(t *testing.T) {
	tests := []struct {
		a int
		b int
	}{
		{a: 1 << 31, b: 1},
		{a: 0, b: 1},
		{a: 1, b: 1},
		{a: -1, b: 1},
		{a: -1, b: -1},
		{a: 10, b: 3},
		{a: -10, b: -3},
		{a: 10, b: -3},
		{a: -10, b: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			want := tt.a / tt.b
			if got := divide(tt.a, tt.b); got != want {
				t.Errorf("divide() = %v, want %v", got, want)
			}
		})
	}
}
