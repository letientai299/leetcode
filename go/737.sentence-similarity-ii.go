package main

// Sentence Similarity II
//
// # Medium
//
// https://leetcode.com/problems/sentence-similarity-ii
func areSentencesSimilarTwo(ss1 []string, ss2 []string, similarPairs [][]string) bool {
	if len(ss1) != len(ss2) {
		return false
	}

	m := make(map[string]map[string]struct{})
	for _, p := range similarPairs {
		a, b := p[0], p[1]
		set, ok := m[a]
		if !ok {
			set = make(map[string]struct{})
			set[a] = struct{}{}
		}

		// merge
		for v := range m[b] {
			set[v] = struct{}{}
			m[v] = set
		}

		set[b] = struct{}{}
		m[a] = set
		m[b] = set
	}

	for i, a := range ss1 {
		b := ss2[i]
		_, ok := m[a][b]
		if a != b && !ok {
			return false
		}
	}

	return true
}
