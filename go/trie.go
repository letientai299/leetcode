package main

import (
	"strings"
)

const (
	englishLowerSize = 26
	trieNodePrefix1  = "└─"
	trieNodePrefix2  = "├─"
	trieNodePrefix3  = "│ "
)

func newTrie() *trieNode {
	return &trieNode{
	}
}

// trieNode support lower case english letters only.
type trieNode struct {
	children  [englishLowerSize]*trieNode
	isWordEnd bool
}

func (t *trieNode) insert(w string) {
	if len(w) == 0 {
		t.isWordEnd = true
		return
	}

	c := w[0] - 'a'
	if t.children[c] == nil {
		t.children[c] = &trieNode{}
	}
	t.children[c].insert(w[1:])
}

func (t trieNode) findPrefix(w string) (string, bool) {
	return t.findPrefixBuf(w, make([]byte, 0, len(w)))
}

func (t trieNode) findPrefixBuf(w string, buf []byte) (string, bool) {
	if t.isWordEnd {
		return string(buf), true
	}

	if len(w) == 0 {
		return "", false
	}

	c := w[0] - 'a'
	if c < 0 || int(c) >= len(t.children) || t.children[c] == nil {
		return "", false
	}

	return t.children[c].findPrefixBuf(w[1:], append(buf, w[0]))
}

func (t trieNode) str(buf []byte) []string {
	var ss []string
	l := len(buf)
	for i, v := range t.children {
		if v == nil {
			continue
		}

		c := byte(i + 'a')
		buf = append(buf, c)
		if v.isWordEnd {
			ss = append(ss, string(c)+"  ("+string(buf)+")")
		} else {
			ss = append(ss, string(c))
		}

		sub := v.str(buf)
		buf = buf[:l]

		if len(sub) == 0 {
			continue
		}

		ss = append(ss, trieNodePrefix1+sub[0])

		for i, s := range sub[1:] {
			if 'a' <= s[0] && s[0] <= 'z' {
				if i >= len(sub)-1 || sub[i+1][0] != ' ' {
					ss = append(ss, trieNodePrefix1+s)
				} else {
					ss = append(ss, trieNodePrefix2+s)
				}
				continue
			}

			ss = append(ss, "  "+s)
		}
	}

	return ss
}

func (t trieNode) String() string {
	sub := t.str([]byte{})

	var ss []string
	for _, s := range sub {
		if 'a' <= s[0] && s[0] <= 'z' {
			ss = append(ss, trieNodePrefix1+s)
			continue
		}

		ss = append(ss, "  "+s)
	}
	return strings.Join(ss, "\n")
}
