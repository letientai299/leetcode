package main

import (
	"fmt"
)

// medium

func fractionToDecimal(N int, D int) string {
	if N == 0 {
		return "0"
	}

	neg := (N >= 0) != (D >= 0)
	n := int64(N)
	d := int64(D)
	if n < 0 {
		n = -n
	}
	if d < 0 {
		d = -d
	}

	buf := make([]byte, 0, 100)
	ns := make(map[int64]int)
	if neg {
		buf = append(buf, '-')
	}
	buf = append(buf, fmt.Sprint(n/d)...)
	buf = append(buf, '.')
	n = (n % d) * 10
	for n > 0 {
		if ns[n] == 0 {
			c := byte(n/d + '0')
			buf = append(buf, c)
			ns[n] = len(buf)-1
			n %= d
			n *= 10
			continue
		}
		i := ns[n]
		rep := string(buf[i:])
		buf = buf[:len(buf)-len(rep)]
		return string(buf) + "(" + rep + ")"
	}
	if buf[len(buf)-1] == '.' {
		buf = buf[:len(buf)-1]
	}
	return string(buf)
}
