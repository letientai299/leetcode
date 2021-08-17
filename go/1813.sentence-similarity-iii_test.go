package main

import "testing"

func Test_areSentencesSimilar(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{
			a:    "C",
			b:    "CB B C",
			want: true,
		},

		{
			a:    "xD iP tqchblXgqvNVdi",
			b:    "FmtdCzv Gp YZf UYJ xD iP tqchblXgqvNVdi",
			want: true,
		},

		{
			a:    "A A AAa",
			b:    "A AAa",
			want: true,
		},

		{
			a:    "aaaa",
			b:    "aaaa dfas",
			want: true,
		},

		{
			a:    "Jane",
			b:    "Jane",
			want: true,
		},

		{
			a:    "Jane",
			b:    "Janeee",
			want: false,
		},

		{
			a:    "My Jane",
			b:    "My name is Jane",
			want: true,
		},

		{
			a:    "My name is Jane",
			b:    "My Jane",
			want: true,
		},

		{
			a:    "My name is Jane",
			b:    "My",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areSentencesSimilar(tt.a, tt.b); got != tt.want {
				t.Errorf("areSentencesSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
