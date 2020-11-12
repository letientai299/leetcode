package main

import "strconv"

func toHexspeak(s string) string {
	n, _ := strconv.Atoi(s)
	var bs []byte
	for n > 0 {
		b := n % 16
		switch {
		case b == 0:
			bs = append(bs, 'O')
		case b == 1:
			bs = append(bs, 'I')
		case b < 10:
			return "ERROR"
		default:
			bs = append(bs, byte(b-10+'A'))
		}
		n /= 16
	}

	l := len(bs)
	for i := 0; i < l/2; i++ {
		bs[i], bs[l-1-i] = bs[l-1-i], bs[i]
	}

	return string(bs)
}
