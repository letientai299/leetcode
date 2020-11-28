package main

import "fmt"

// medium

func complexNumberMultiply(a string, b string) string {
	parse := func(s string) (real, magic int) {
		sig := 1
		v := 0
		for _, x := range s {
			switch x {
			case '+':
				real = v * sig
				sig = 1
				v = 0
			case '-':
				sig = -1
				v = 0
			case 'i':
				magic = v * sig
			default:
				v = 10*v + int(x-'0')
			}
		}
		return
	}

	m, n := parse(a)
	p, q := parse(b)

	x, y := m*p-n*q, m*q+n*p

	return fmt.Sprintf("%d+%di", x, y)
}
