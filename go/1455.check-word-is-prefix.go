package main

import "strings"

func isPrefixOfWord(sentence string, prefix string) int {
	ss := strings.Split(sentence, " ")
	for i, s := range ss {
		if strings.HasPrefix(s, prefix) {
			return i + 1
		}
	}
	return -1
}
