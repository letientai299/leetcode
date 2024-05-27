package main

import (
	"strconv"
)

// Unique Word Abbreviation
//
// # Medium
//
// https://leetcode.com/problems/unique-word-abbreviation/
type ValidWordAbbr struct {
	m    map[string]int
	dict map[string]bool
}

func Constructor288(dict []string) ValidWordAbbr {
	vw := ValidWordAbbr{
		m:    make(map[string]int),
		dict: make(map[string]bool),
	}

	for _, w := range dict {
		if vw.dict[w] {
			continue
		}

		vw.dict[w] = true
		vw.m[vw.abbr(w)]++
	}

	return vw
}

func (vw *ValidWordAbbr) abbr(w string) string {
	if len(w) <= 2 {
		return w
	}
	return w[0:1] + strconv.Itoa(len(w)-2) + w[len(w)-1:]
}

func (vw *ValidWordAbbr) IsUnique(w string) bool {
	ab := vw.abbr(w)
	return vw.m[ab] == 0 || (vw.dict[w] && vw.m[ab] == 1)
}
