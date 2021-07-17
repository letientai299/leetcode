package main

import "testing"

func Test_isSumEqual(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		c    string
		want bool
	}{
		{a: "acb", b: "cba", c: "cdb", want: true},
		{a: "aaa", b: "a", c: "aab", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSumEqual(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("isSumEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
