package main

import (
	"testing"
)

func Test_areSentencesSimilar(t *testing.T) {
	tests := []struct {
		a     []string
		b     []string
		pairs [][]string
		want  bool
	}{
		{

			a: []string{"an", "extraordinary", "meal"},
			b: []string{"one", "good", "dinner"},
			pairs: [][]string{
				{"great", "good"},
				{"extraordinary", "good"},
				{"well", "good"},
				{"wonderful", "good"},
				{"excellent", "good"},
				{"fine", "good"},
				{"nice", "good"},
				{"any", "one"},
				{"some", "one"},
				{"unique", "one"},
				{"the", "one"},
				{"an", "one"},
				{"single", "one"},
				{"a", "one"},
				{"truck", "car"},
				{"wagon", "car"},
				{"automobile", "car"},
				{"auto", "car"},
				{"vehicle", "car"},
				{"entertain", "have"},
				{"drink", "have"},
				{"eat", "have"},
				{"take", "have"},
				{"fruits", "meal"},
				{"brunch", "meal"},
				{"breakfast", "meal"},
				{"food", "meal"},
				{"dinner", "meal"},
				{"super", "meal"},
				{"lunch", "meal"},
				{"possess", "own"},
				{"keep", "own"},
				{"have", "own"},
				{"extremely", "very"},
				{"actually", "very"},
				{"really", "very"},
				{"super", "very"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := areSentencesSimilar(tt.a, tt.b, tt.pairs); got != tt.want {
				t.Errorf("areSentencesSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
