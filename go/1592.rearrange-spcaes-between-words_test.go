package main

import "testing"

func Test_reorderSpaces(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{
			text: "  this",
			want: "this  ",
		},

		{
			text: "  this   is  a sentence ",
			want: "this   is   a   sentence",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := reorderSpaces(tt.text); got != tt.want {
				t.Errorf("reorderSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
