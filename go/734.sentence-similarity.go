package main

// https://leetcode.com/problems/sentence-similarity/
func areSentencesSimilar734(a []string, b []string, pairs [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[string]map[string]struct{}, len(pairs)*2)
	for _, ss := range pairs {
		x, y := ss[0], ss[1]

		if list, ok := m[x]; ok {
			list[y] = struct{}{}
		} else {
			m[x] = map[string]struct{}{y: {}}
		}

		if list, ok := m[y]; ok {
			list[x] = struct{}{}
		} else {
			m[y] = map[string]struct{}{x: {}}
		}
	}

	for i, sa := range a {
		sb := b[i]
		if sb == sa {
			continue
		}

		p, ok := m[sa]
		if !ok {
			return false
		}

		if _, ok := p[sb]; !ok {
			return false
		}
	}

	return true
}
