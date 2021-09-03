package main

import "testing"

func Test_equationsPossible(t *testing.T) {
	tests := []struct {
		name      string
		equations []string
		want      bool
	}{
		{equations: []string{"a==b", "b!=a"}, want: false},
		{equations: []string{"b==a", "a==b"}, want: true},
		{equations: []string{"a==b", "b==c", "a==c"}, want: true},
		{equations: []string{"a==b", "b!=c", "c==a"}, want: false},
		{equations: []string{"c==c", "b==d", "x!=z"}, want: true},
		{equations: []string{"a!=a"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
