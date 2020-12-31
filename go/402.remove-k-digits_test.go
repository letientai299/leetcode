package main

import "testing"

func Test_removeKdigits(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{s: "5337", k: 2, want: "33"},
		{s: "10", k: 0, want: "10"},
		{s: "10", k: 1, want: "0"},
		{s: "10", k: 2, want: "0"},
		{s: "10200", k: 1, want: "200"},
		{s: "10000000234444121", k: 3, want: "2344121"},
		{s: "1000000023456789", k: 3, want: "234567"},
		{s: "123456789", k: 3, want: "123456"},
		{s: "123456001", k: 3, want: "123001"},
		{s: "1439919", k: 3, want: "1319"},
		{s: "1432219", k: 3, want: "1219"},
		{s: "123456001", k: 1, want: "12345001"},
		{s: "123456001", k: 0, want: "123456001"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := removeKdigits(tt.s, tt.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
