package main

import (
	"fmt"
	"testing"
)

func Test_getHint(t *testing.T) {
	tests := []struct {
		secret string
		guess  string
		want   string
	}{
		{secret: "1807", guess: "7810", want: "1A3B"},
		{secret: "1123", guess: "0111", want: "1A1B"},
		{secret: "11", guess: "10", want: "1A0B"},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.secret, tt.guess,
		)
		t.Run(testName, func(t *testing.T) {
			got := getHint(tt.secret, tt.guess)
			if got != tt.want {
				t.Errorf("getHint(%v, %v) = %v, want %v", tt.secret, tt.guess, got, tt.want)
			}
		})
	}
}
