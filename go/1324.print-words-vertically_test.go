package main

import (
	"reflect"
	"testing"
)

func Test_printVertically(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{s: "HOW ARE YOU", want: []string{"HAY", "ORO", "WEU"}},
		{s: "TO BE OR NOT TO BE", want: []string{"TBONTB", "OEROOE", "   T"}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := printVertically(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printVertically() = %v, want %v", got, tt.want)
			}
		})
	}
}
