package main

import (
	"testing"
)

func Test_confusingNumber(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{n: 96, want: false},
		{n: 6, want: true},
		{n: 11, want: false},
		{n: 12, want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := confusingNumber(tt.n); got != tt.want {
				t.Errorf("confusingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
