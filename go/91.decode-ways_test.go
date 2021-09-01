package main

import "testing"

func Test_numDecodings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		// 2 6 1 10 5 5 9 7 1 7 5 6 5 6 2
		// 2 6 1 10 5 5 9 7 17 5 6 5 6 2
		// 26 1 10 5 5 9 7 17 5 6 5 6 2
		// 26 1 10 5 5 9 7 1 7 5 6 5 6 2
		{s: "2611055971756562", want: 4},
		{s: "227", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := numDecodings(tt.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
