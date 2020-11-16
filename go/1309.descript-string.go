package main

import "strings"

func freqAlphabets(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		x := s[i]
		if x == '#' {
			c := (s[i-2]-'0')*10 + (s[i-1] - '0') + 'a'
			sb.WriteByte(c)
			i -= 2
			continue
		}
		sb.WriteByte('a' - '1' + x)
	}
	b := []byte(sb.String())
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b)
}
