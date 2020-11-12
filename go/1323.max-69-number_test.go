package main

import "testing"

func Test_maximum69Number(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{num: -69, want: -69},
		{num: 69, want: 99},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximum69Number(tt.num); got != tt.want {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
