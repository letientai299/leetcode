package main

import (
	"reflect"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			s: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			want: []string{
				"AAAAACCCCC", "CCCCCAAAAA",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := findRepeatedDnaSequences(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
