package main

import "testing"

func Test_arrangeWords(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{text: "Leetcode is cool", want: "Is cool leetcode"},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			if got := arrangeWords(tt.text); got != tt.want {
				t.Errorf("arrangeWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
