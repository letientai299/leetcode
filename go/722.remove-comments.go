package main

import "strings"

/*
 * @lc app=leetcode id=722 lang=golang
 *
 * [722] Remove Comments
 *
 * https://leetcode.com/problems/remove-comments/description/
 */

// Copied from https://leetcode.com/problems/remove-comments/discuss/979346/golang-with-strings.Builder
// because this problem is not fun to play with.
func removeComments(source []string) []string {
	output := make([]string, 0)

	insideBlock := false
	sb := strings.Builder{}
	tmpRune := ' '
	lnLen := 0
	var data []rune

	for _, line := range source {
		lnLen = len(line)
		if !insideBlock {
			sb.Reset()
			sb.Grow(lnLen)
		}
		data = []rune(line)
		for i := 0; i < lnLen; i++ {
			tmpRune = data[i]
			if !insideBlock {
				if tmpRune != '/' {
					sb.WriteRune(tmpRune)
				} else {
					if lnLen > i+1 {
						if data[i+1] == '*' {
							insideBlock = true
							i++
							continue
						} else if data[i+1] == '/' {
							break
						} else {
							sb.WriteRune(tmpRune)
						}
					} else {
						sb.WriteRune(tmpRune)
					}
				}
			} else if tmpRune == '*' && data[i+1] == '/' {
				i++
				insideBlock = false
			}
		}
		if !insideBlock && sb.Len() > 0 {
			output = append(output, sb.String())
		}
	}
	return output
}
