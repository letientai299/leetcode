package main

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func max(a int, arr ...int) int {
	for _, v := range arr {
		if a < v {
			a = v
		}
	}

	return a
}

func min(a int, arr ...int) int {
	for _, v := range arr {
		if a > v {
			a = v
		}
	}

	return a
}

// binary gcd
func gcd(a, b int) int {
	if a == b {
		return a
	}

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	u, v := uint(a), uint(b)
	shift := trailing0bit(u | v)
	u >>= trailing0bit(u)
	for v != 0 {
		v >>= trailing0bit(v)
		if u > v {
			u, v = v, u
		}
		v = v - u
	}

	return int(u << shift)
}

func trailing0bit(x uint) uint {
	n := uint(0)
	for x > 0 && x%2 == 0 {
		n += 1 - (x % 2)
		x /= 2
	}
	return n
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func numPalindrome(n int) bool {
	v := n
	r := 0
	for v > 0 {
		r = 10*r + v%10
		v /= 10
	}
	return r == n
}

func primeFactor(n int) []int {
	var r []int
	i := 2
	if n < i {
		return r
	}

	for n >= i {
		if n%i == 0 {
			r = append(r, i)
			for n%i == 0 {
				n /= i
			}
		}
		if i > 2 {
			i += 2
		} else {
			i++
		}
	}

	return r
}

func isPrime(a int) bool {
	n := a
	if n%2 == 0 {
		return false
	}

	if n == 3 || n == 5 || n == 7 || n == 11 {
		return true
	}

	for i := 3; i <= int(math.Floor(math.Sqrt(float64(n)))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func divisors(n int) []int {
	if n <= 0 {
		return []int{1}
	}

	var r []int
	i := 1
	for ; i < int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			r = append(r, i, n/i)
		}
	}

	if n%i == 0 {
		r = append(r, i)
		if n/i != i {
			r = append(r, n/i)
		}
	}

	return r
}

func sum(arr ...int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func quickSelect(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	pivot := nums[int(rand.Uint32()%uint32(n))]
	c := 0
	l, r := 0, n-1
	for l <= r {
		for l <= r && nums[l] <= pivot {
			if nums[l] == pivot {
				c++
			}
			l++
		}

		for r >= l && nums[r] > pivot {
			r--
		}

		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	if l >= k && k > l-c {
		return pivot
	}

	if l-c >= k {
		return quickSelect(nums[:l], k)
	}

	return quickSelect(nums[l:], k-l)
}

func quickSelectFunc(n, k int,
	f func(i int) interface{},
	comp func(a, b interface{}) int,
	swap func(i, j int),
) interface{} {
	if n == 1 {
		return f(0)
	}

	pi := int(rand.Uint32() % uint32(n))
	pivot := f(pi)

	c := 0
	l, r := 0, n-1
	for l <= r {
		for l <= r {
			a := f(l)
			cmp := comp(a, pivot)
			if cmp > 0 {
				break
			}
			if cmp == 0 {
				c++
			}
			l++
		}

		for b := f(r); r >= l && comp(b, pivot) > 0; {
			r--
			b = f(r)
		}

		if l < r {
			swap(l, r)
			l++
			r--
		}
	}

	if l >= k && k > l-c {
		return pivot
	}

	if l-c >= k {
		return quickSelectFunc(l, k,
			func(i int) interface{} { return f(i) },
			comp,
			swap,
		)
	}

	return quickSelectFunc(n-l, k-l,
		func(i int) interface{} { return f(i + l) },
		comp,
		swap,
	)
}
