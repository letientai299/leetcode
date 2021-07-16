package main

import "testing"

func Test_secondHighest(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "sjhtz8344", want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := secondHighest(tt.s); got != tt.want {
				t.Errorf("secondHighest() = %v, want %v", got, tt.want)
			}
		})
	}
}
