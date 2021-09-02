package main

import "testing"

func Test_getHappyString(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want string
	}{
		{n: 3, k: 9, want: "cab"},
		{n: 3, k: 10, want: "cac"},
		{n: 3, k: 11, want: "cba"},
		{n: 3, k: 12, want: "cbc"},
		{n: 3, k: 13, want: ""},
		{n: 1, k: 1, want: "a"},
		{n: 1, k: 2, want: "b"},
		{n: 1, k: 3, want: "c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHappyString(tt.n, tt.k); got != tt.want {
				t.Errorf("getHappyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
