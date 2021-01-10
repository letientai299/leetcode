package main

import "testing"

func Test_getSmallestString(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want string
	}{
		{n: 5, k: 73, want: "aaszz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSmallestString(tt.n, tt.k); got != tt.want {
				t.Errorf("getSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
