package main

// Shortest Word Distance II
//
// # Medium
//
// https://leetcode.com/problems/shortest-word-distance-ii/
type WordDistance struct {
	m   map[string][]int
	max int
}

func Constructor244(wordsDict []string) WordDistance {
	d := WordDistance{
		m:   make(map[string][]int),
		max: len(wordsDict),
	}
	for i, w := range wordsDict {
		d.m[w] = append(d.m[w], i)
	}
	return d
}

func (d *WordDistance) Shortest(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}

	a, b := d.m[word1], d.m[word2]
	i, j := 0, 0
	best := d.max
	for i < len(a) && j < len(b) {
		best = min(best, abs(a[i]-b[j]))
		if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}

	return best
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */
