package main

import "testing"

func Test_maxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{text: "balloon", want: 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxNumberOfBalloons(tt.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", got, tt.want)
			}
		})
	}
}
