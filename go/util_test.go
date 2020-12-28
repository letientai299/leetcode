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
		{a: 3, b: 8, want: 1},
		{a: 9, b: 12, want: 3},
		{a: 12, b: 9, want: 3},
		{a: 4, b: 7, want: 1},
		{a: 4, b: 8, want: 4},
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
		{name: "gcdEuler", fn: gcdEuler},
		{name: "gcdEulerRecursive", fn: gcdEulerRecursive},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tt.fn(2147483647, 24036583)
				tt.fn(1239203, 21321824)
			}
		})
	}
}

func gcdEuler(a, b int) int {
	for a != 0 && a != 1 {
		a, b = b%a, a
	}

	if a == 0 {
		return b
	}

	return 1
}

func gcdEulerRecursive(a, b int) int {
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

func Test_divisors(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{n: 0, want: []int{1}},
		{n: 3, want: []int{1, 3}},
		{n: 4, want: []int{1, 2, 4}},
		{n: 9, want: []int{1, 3, 9}},
		{n: 16, want: []int{1, 2, 4, 8, 16}},
		{n: 21, want: []int{1, 21, 3, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisors(tt.n); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("divisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trailing0bit(t *testing.T) {
	tests := []struct {
		name string
		x    uint
		want uint
	}{
		{x: 11, want: 0},
		{x: 0, want: 0},
		{x: 1, want: 0},
		{x: 2, want: 1},
		{x: 3, want: 0},
		{x: 4, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailing0bit(tt.x); got != tt.want {
				t.Errorf("trailing0bit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSelect(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 3, 2}, k: 2, want: 2},
		{nums: []int{1, 2, 3, 4, 5}, k: 2, want: 2},
		{nums: []int{1, 2, 3, 4, 5}, k: 4, want: 4},
		{nums: []int{1, 2, 3, 4, 5}, k: 5, want: 5},
		{nums: []int{1, 2, 3, 4, 5}, k: 1, want: 1},
		{nums: []int{1, 2, 3, 4, 5, 6}, k: 1, want: 1},
		{nums: []int{1, 2, 3, 4, 5, 6}, k: 3, want: 3},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			qs := quickSelect(tt.nums, tt.k)
			if qs != tt.want {
				t.Errorf("quickSelect() = %v, want %v", qs, tt.want)
				return
			}

			find := func(i int) interface{} { return tt.nums[i] }
			swap := func(i, j int) { tt.nums[i], tt.nums[j] = tt.nums[j], tt.nums[i] }
			comp := func(a, b interface{}) int {
				x := a.(int)
				y := b.(int)
				if x == y {
					return 0
				}

				if x < y {
					return -1
				}
				return 1
			}
			qsf := quickSelectFunc(len(tt.nums), tt.k, find, comp, swap)
			if qsf != tt.want {
				t.Errorf("quickSelectFunc() = %v, want %v", qsf, tt.want)
			}
		})
	}
}
