package main

import (
	"strings"
)

// https://leetcode.com/problems/bold-words-in-string

func boldWords(words []string, s string) string {
	mask := make([]bool, len(s))
	for _, w := range words {
		for i, skip := strings.Index(s, w), 0; i >= 0; {
			for x := 0; x < len(w); x++ {
				mask[x+skip+i] = true
			}
			skip = i + skip + 1
			i = strings.Index(s[skip:], w)
		}
	}
	var sb strings.Builder
	prev := false
	for i, c := range s {
		if mask[i] != prev {
			if mask[i] {
				sb.Write([]byte("<b>"))
			} else {
				sb.Write([]byte("</b>"))
			}
		}
		sb.WriteByte(byte(c))
		prev = mask[i]
	}

	if prev {
		sb.Write([]byte("</b>"))
	}

	return sb.String()
}
