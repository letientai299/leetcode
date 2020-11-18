package main

import "strings"

func reorderSpaces(text string) string {
	cs := strings.Count(text, " ")
	buf := make([]byte, 0, len(text))
	ws := strings.Split(text, " ")

	cw := 0
	for {
		for cw < len(ws) && ws[cw] != "" {
			cw++
		}
		j := cw + 1
		for j < len(ws) && ws[j] == "" {
			j++
		}
		if j < len(ws) {
			ws[cw], ws[j] = ws[j], ws[cw]
		} else {
			break
		}
	}
	ws = ws[:cw]
	ns := 0
	if cw > 1 {
		ns = cs / (cw - 1)
		cs %= cw - 1
	}
	for i, w := range ws {
		buf = append(buf, w...)
		if i == cw-1 {
			break
		}
		for j := 0; j < ns; j++ {
			buf = append(buf, ' ')
		}
	}

	for j := 0; j < cs; j++ {
		buf = append(buf, ' ')
	}
	return string(buf)
}
