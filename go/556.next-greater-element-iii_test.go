package main

import "testing"

func Test_nextGreaterElement(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 12443322, want: 13222344},
		{n: 234567775, want: 234575677},
		{n: 456775321, want: 457123567},
		{n: 1, want: -1},
		{n: 123456775321, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.n); got != tt.want {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
