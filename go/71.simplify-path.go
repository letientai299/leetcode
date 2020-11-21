package main

import "strings"

// medium
// https://leetcode.com/problems/simplify-path/

func simplifyPath(path string) string {
	var st []string
	ss := strings.Split(path, "/")
	for _, s := range ss {
		switch s {
		case "", ".":
			continue
		case "..":
			if len(st) != 0 {
				st = st[:len(st)-1]
			}
		default:
			st = append(st, s)
		}
	}

	return "/" + strings.Join(st, "/")
}
