package main

import (
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	var res []string
	for i, w := range words {
		for j := len(words) - 1; j > i; j-- {
			if strings.Contains(words[j], w) {
				res = append(res, w)
				break
			}
		}
	}
	return res
}
