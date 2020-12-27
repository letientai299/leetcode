package main

import "testing"

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
			want:  16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.words); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
