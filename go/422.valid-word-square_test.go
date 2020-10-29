package main

import (
	"fmt"
	"testing"
)

func Test_validWordSquare(t *testing.T) {
	tests := []struct {
		words []string
		want  bool
	}{
		{
			words: []string{
				"abc",
				"b",
			},
			want: false,
		},

		{
			words: []string{
				"abcd",
				"bnrt",
				"crm",
				"dt",
			},
			want: true,
		},

		{
			words: []string{
				"ball",
				"asee",
				"let",
				"lep",
			},
			want: false,
		},

		{
			words: []string{
				"b1cd",
				"bnrt",
				"crmy",
				"dtye",
			},
			want: false,
		},
		{
			words: []string{
				"abcd",
				"bnrt",
				"crmy",
				"dtye",
			},
			want: true,
		},

		{
			words: []string{
				"abc",
				"bef",
				"cfg",
			},
			want: true,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.words,
		)
		t.Run(testName, func(t *testing.T) {
			got := validWordSquare(tt.words)
			if got != tt.want {
				t.Errorf("validWordSquare(%v) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}
