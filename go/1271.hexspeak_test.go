package main

import "testing"

func Test_toHexspeak(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "256", want: "IOO"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := toHexspeak(tt.s); got != tt.want {
				t.Errorf("toHexspeak() = %v, want %v", got, tt.want)
			}
		})
	}
}
