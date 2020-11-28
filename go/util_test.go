package main

import "testing"

func Test_gcd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{a: 4, b: 7, want: 1},
		{a: 4, b: 8, want: 4},
		{a: 3, b: 8, want: 1},
		{a: 9, b: 12, want: 3},
		{a: 12, b: 9, want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := gcd(tt.a, tt.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_gcd(b *testing.B) {
	tests := []struct {
		name string
		fn   func(a, b int) int
	}{
		{name: "gcd", fn: gcd},
		{name: "gcdRecursive", fn: gcdRecursive},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tt.fn(2147483647, 24036583)
			}
		})
	}
}
func gcdRecursive(a, b int) int {
	if a == 1 {
		return 1
	}
	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}
