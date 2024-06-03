package main

import (
	"sort"
	"strings"
)

// Add Bold Tag in String
//
// # Medium
//
// https://leetcode.com/problems/add-bold-tag-in-string
func addBoldTag(s string, words []string) string {
	type pair struct {
		x, y int
	}

	var ps []pair
	for _, w := range words {
		start := 0
		for start < len(s) {
			x := strings.Index(s[start:], w)
			if x < 0 {
				break
			}
			x += start
			ps = append(ps, pair{x: x, y: x + len(w)})
			start = x + 1
		}
	}

	if len(ps) == 0 {
		return s
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].x < ps[j].x
	})

	var sb strings.Builder
	cur := ps[0]
	if cur.x > 0 {
		sb.WriteString(s[:cur.x])
	}

	for _, p := range ps {
		if cur.y < p.x {
			sb.WriteString("<b>")
			sb.WriteString(s[cur.x:cur.y])
			sb.WriteString("</b>")
			sb.WriteString(s[cur.y:p.x])
			cur = p
			continue
		}

		cur.y = max(cur.y, p.y)
	}

	sb.WriteString("<b>")
	sb.WriteString(s[cur.x:cur.y])
	sb.WriteString("</b>")
	sb.WriteString(s[cur.y:])
	return sb.String()
}
