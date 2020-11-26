package main

import "testing"

func Test_integerReplacement(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 26, want: 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := integerReplacement(tt.n); got != tt.want {
				t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
