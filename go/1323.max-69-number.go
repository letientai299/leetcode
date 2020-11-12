package main

func maximum69Number(num int) int {
	n := num
	max := n
	f := 1
	sign := 1
	if num < 0 {
		sign = -1
	}
	for n > 0 {
		d := n % 10
		other := num - sign*f*d
		if d == 6 {
			other += sign * f * 9
		} else {
			other += sign * f * 6
		}
		if other > max {
			max = other
		}
		f *= 10
		n /= 10
	}
	return max
}
