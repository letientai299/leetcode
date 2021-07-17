package main

import "testing"

func Test_makeEqual(t *testing.T) {
	tests := []struct {
		words []string
		want  bool
	}{
		{words: []string{"abc", "aabc", "bc"}, want: true},
		{words: []string{"ab", "a", "bc"}, want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := makeEqual(tt.words); got != tt.want {
				t.Errorf("makeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
