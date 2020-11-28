package main

import "fmt"

// medium

func solveEquation(equation string) string {
	x, v := 0, 0
	n := 0
	hasDigit := false
	sig := 1
	side := 1 // -1 mean right side
	for _, c := range equation {
		switch c {
		case '=':
			hasDigit = false
			v += n * sig * side
			n = 0
			sig = 1
			side = -1
		case 'x':
			if n == 0 && !hasDigit {
				n = 1
			}
			x += n * sig * side
			n = 0
			hasDigit = false
		case '+':
			v += n * sig * side
			n = 0
			sig = 1
			hasDigit = false
		case '-':
			v += n * sig * side
			sig = -1
			n = 0
			hasDigit = false
		default:
			n = 10*n + int(c-'0')
			hasDigit = true
		}
	}

	v += n * sig * side

	if x == 0 {
		if v == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}

	return fmt.Sprintf("x=%d", -v/x)
}
