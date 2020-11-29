package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func Test_numPalindrome(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{n: 12, want: false},
		{n: 10001, want: true},
		{n: 12321, want: true},
		{n: 121, want: true},
		{n: 11, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPalindrome(tt.n); got != tt.want {
				t.Errorf("numPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fromDigits(t *testing.T) {
	ds := digits(0)
	assert.ElementsMatch(t, ds, []int{0})
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 10; i++ {
		n := rand.Int()
		ds = digits(n)
		if got := fromDigits(ds); got != n {
			t.Errorf("fromDigits() = %v, want %v", got, n)
		}
	}
}

func Test_primeFactor(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{n: 2, want: []int{2}},
		{n: 12, want: []int{2, 3}},
		{n: 2 * 3 * 5 * 7 * 11, want: []int{2, 3, 5, 7, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primeFactor(tt.n); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("primeFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
