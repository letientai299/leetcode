package main

import "strings"

// medium

func originalDigits(s string) string {
	m := make(map[byte]int)
	for _, c := range s {
		m[byte(c)]++
	}
	// zero one two three four five six seven eight nine
	ds := [10]int{}

	if m['x'] != 0 {
		m['s'] -= m['x']
		m['i'] -= m['x']
		ds[6] = m['x']
	}

	if m['w'] != 0 {
		m['t'] -= m['w']
		ds[2] = m['w']
		m['o'] -= m['w']
	}

	if m['g'] != 0 {
		m['e'] -= m['g']
		m['i'] -= m['g']
		ds[8] = m['g']
		m['h'] -= m['g']
		m['t'] -= m['g']
	}

	if m['u'] != 0 {
		ds[4] = m['u']
		m['f'] -= m['u']
		m['o'] -= m['u']
		m['r'] -= m['u']
	}

	if m['z'] != 0 {
		ds[0] = m['z']
		m['e'] -= m['z']
		m['r'] -= m['z']
		m['o'] -= m['z']
	}

	// one three five seven nine

	if m['o'] != 0 {
		ds[1] = m['o']
		m['n'] -= m['o']
		m['e'] -= m['o']
	}

	if m['t'] != 0 {
		ds[3] = m['t']
		m['h'] -= m['t']
		m['r'] -= m['t']
		m['e'] -= 2 * m['t']
	}

	if m['f'] != 0 {
		ds[5] = m['f']
		m['i'] -= m['f']
		m['v'] -= m['f']
		m['e'] -= m['f']
	}

	if m['s'] != 0 {
		ds[7] = m['s']
		m['e'] -= 2 * m['s']
		m['v'] -= m['s']
		m['n'] -= m['s']
	}

	if m['i'] != 0 {
		ds[9] = m['i']
	}

	var sb strings.Builder
	for i, n := range ds {
		for j := 0; j < n; j++ {
			sb.WriteByte(byte(i + '0'))
		}
	}
	return sb.String()
}
