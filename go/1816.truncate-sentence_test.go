package main

import "testing"

func Test_truncateSentence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{s: "Hello how are you Contestant", k: 4, want: "Hello how are you"},
		{s: "a b c d e", k: 4, want: "a b c d"},
		{s: "hello there", k: 1, want: "hello"},
		{s: "hello there", k: 2, want: "hello there"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncateSentence(tt.s, tt.k); got != tt.want {
				t.Errorf("truncateSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
