package main

import "testing"

func Test_canConvertString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		k    int
		want bool
	}{
		{s: "abc", t: "bcd", k: 10, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConvertString(tt.s, tt.t, tt.k); got != tt.want {
				t.Errorf("canConvertString() = %v, want %v", got, tt.want)
			}
		})
	}
}
