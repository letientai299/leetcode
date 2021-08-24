package main

import "testing"

func Test_maxValue(t *testing.T) {
	tests := []struct {
		n    string
		x    int
		want string
	}{
		{n: "43341", x: 3, want: "433431"},
		{n: "123", x: 3, want: "3123"},
		{n: "-132", x: 3, want: "-1323"},
		{n: "-13", x: 2, want: "-123"},
	}
	for _, tt := range tests {
		t.Run(tt.n, func(t *testing.T) {
			if got := maxValue(tt.n, tt.x); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
