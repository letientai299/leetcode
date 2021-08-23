package main

import "testing"

func Test_longestDiverseString(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
		want string
	}{
		{a: 1, b: 1, c: 7, want: "ccaccbcc"},
		{a: 2, b: 2, c: 1, want: "ababc"},
		{a: 7, b: 1, c: 0, want: "aabaa"},
		{a: 0, b: 0, c: 0, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDiverseString(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("longestDiverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
