package main

import "testing"

func Test_expressiveWords(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  int
	}{
		{
			s:     "lee",
			words: []string{"le"},
			want:  0,
		},

		{
			s:     "heeellooo",
			words: []string{"hello", "hi", "helo"},
			want:  1,
		},


		{
			s:     "sass",
			words: []string{"sa"},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expressiveWords(tt.s, tt.words); got != tt.want {
				t.Errorf("expressiveWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
