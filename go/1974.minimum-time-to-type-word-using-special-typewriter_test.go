package main

import "testing"

func Test_minTimeToType(t *testing.T) {
	tests := []struct {
		word string
		want int
	}{
		{word: "abc", want: 5},
		{word: "bza", want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := minTimeToType(tt.word); got != tt.want {
				t.Errorf("minTimeToType() = %v, want %v", got, tt.want)
			}
		})
	}
}
